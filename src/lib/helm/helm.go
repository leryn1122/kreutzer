package helm

import (
	"context"
	"github.com/gofrs/flock"
	"github.com/leryn1122/kreutzer/v2/kreutzer/dao"
	"github.com/leryn1122/kreutzer/v2/kreutzer/vo"
	"github.com/leryn1122/kreutzer/v2/lib/db"
	"github.com/leryn1122/kreutzer/v2/support/path"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/kube"
	chartrepo "helm.sh/helm/v3/pkg/repo"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
const (
	HelmDriverKey                 = "HELM_DRIVER"
	HelmRepositoryConfigKey       = "HELM_REPOSITORY_CONFIG"
	HelmRepositoryCacheKey        = "HELM_REPOSITORY_CACHE"
	HelmRepositoryCacheFileSuffix = "-index.yaml"
)

type Client struct {
	flock      *flock.Flock
	settings   *cli.EnvSettings
	helmDriver string
}

type ClientOption = func(client *Client)

type KubeInfo struct {
	namespace string
	context   string
	config    string
}

type Repo struct {
	Name                  string    `json:"name"`
	URL                   string    `json:"url"`
	LastSyncTime          time.Time `json:"lastSyncTime"`
	Username              string    `json:"username"`
	Password              string    `json:"password"`
	PassCredentialsAll    bool      `json:"passCredentialsAll"`
	InsecureSkipTLSVerify bool      `json:"insecureSkipTLSVerify"`
	CertFile              string    `json:"certFile"`
	KeyFile               string    `json:"keyFile"`
	CAFile                string    `json:"caFile"`
}

type RepoErr struct {
	Name string
	Err  string
}

type Release struct {
	Name       string    `json:"name"`
	Namespace  string    `json:"namespace"`
	Revision   int       `json:"revision"`
	Updated    time.Time `json:"updated"`
	Status     string    `json:"status"`
	Chart      string    `json:"chart"`
	AppVersion string    `json:"appVersion"`
}

func NewHelmClient(opts ...ClientOption) *Client {
	settings := cli.New()
	var client = &Client{
		flock:      nil,
		settings:   settings,
		helmDriver: os.Getenv(HelmDriverKey),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(client)
		}
	}
	return client
}

func (c *Client) acquireRepositoryFileLock() error {
	if c.flock != nil {
		if err := c.createRepositoryFileLock(); err != nil {
			return err
		}
	}

	fileLock := c.flock
	lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	if err == nil && locked {
	}

	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createRepositoryFile() error {
	configFile := c.getRepositoryFile()
	err := pathutil.CreateFileIfNotExists(configFile)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) getRepositoryFile() string {
	path := c.settings.EnvVars()[HelmRepositoryConfigKey]
	return os.ExpandEnv(path)
}

func (c *Client) loadRepositoryFile() (*chartrepo.File, error) {
	path := c.getRepositoryFile()
	return chartrepo.LoadFile(path)
}

func (c *Client) createRepositoryFileLock() error {
	configPath := c.settings.EnvVars()[HelmRepositoryConfigKey]
	dirname := filepath.Dir(configPath)
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	c.flock = flock.New("")

	return nil
}

func (c *Client) createRepositoryCacheFile(repo string) error {
	err := pathutil.CreateFileIfNotExists(c.getRepositoryCacheFile(repo))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) getRepositoryCacheDir() string {
	path := c.settings.EnvVars()[HelmRepositoryCacheKey]
	return os.ExpandEnv(path)
}

func (c *Client) getRepositoryCacheFile(repo string) string {
	return filepath.Join(c.getRepositoryCacheDir(), repo+HelmRepositoryCacheFileSuffix)
}

func (c *Client) loadRepositoryCacheFile(repo string) (map[string]chartrepo.ChartVersions, error) {
	indexFile, err := chartrepo.LoadIndexFile(c.getRepositoryCacheFile(repo))
	if err != nil {
		return nil, err
	}
	return indexFile.Entries, nil
}

func (c *Client) actionConfigInit(kubeInfo *KubeInfo) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)
	if kubeInfo.context == "" {
		kubeInfo.context = c.settings.KubeContext
	}

	var clientConfig *genericclioptions.ConfigFlags
	if kubeInfo.config == "" {
		clientConfig = kube.GetConfig(c.settings.KubeConfig, kubeInfo.context, kubeInfo.namespace)
	} else {
		clientConfig = kube.GetConfig(kubeInfo.config, kubeInfo.context, kubeInfo.namespace)
	}

	if c.settings.KubeToken != "" {
		clientConfig.BearerToken = &c.settings.KubeToken
	}
	if c.settings.KubeAPIServer != "" {
		clientConfig.APIServer = &c.settings.KubeAPIServer
	}

	if err := actionConfig.Init(clientConfig, kubeInfo.namespace, c.helmDriver, logrus.Infof); err != nil {
		logrus.Errorf("%+v", err)
		return nil, err
	}
	return actionConfig, nil
}

// ListRepos
// List repos from database
func (c *Client) ListRepos() ([]*Repo, error) {
	var repos []dao.HelmRepo
	tx := db.DBClient.Find(&repos)
	if err := tx.Error; err != nil {
		return nil, err
	}
	var repos_ []*Repo
	for _, repo_ := range repos {
		repo_ := &Repo{
			Name:                  repo_.Name,
			URL:                   repo_.URL,
			LastSyncTime:          repo_.LastSyncTime,
			Username:              repo_.Username,
			Password:              repo_.Password,
			PassCredentialsAll:    repo_.PassCredentialsAll,
			InsecureSkipTLSVerify: repo_.InsecureSkipTLSVerify,
			CAFile:                repo_.CAFile,
			CertFile:              repo_.CertFile,
			KeyFile:               repo_.KeyFile,
		}
		repos_ = append(repos_, repo_)
	}
	return repos_, nil
}

func (c *Client) AddRepo() error {
	return nil
}

func (c *Client) UpdateRepo(name string) error {
	file, err := c.loadRepositoryFile()
	repo, err := chartrepo.NewChartRepository(file.Get(name), getter.All(c.settings))
	if err != nil {
		return err
	}
	_, err = repo.DownloadIndexFile()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateRepos() []RepoErr {
	var repoErrs []RepoErr
	file, err := c.loadRepositoryFile()
	if err != nil {
	}
	var wg sync.WaitGroup
	for _, repo_ := range file.Repositories {
		wg.Add(1)
		go func(entry *chartrepo.Entry) {
			defer wg.Done()
			err := c.UpdateRepo(entry.Name)
			if err != nil {
				repoErrs = append(repoErrs, RepoErr{
					Name: entry.Name,
					Err:  err.Error(),
				})
			}
		}(repo_)
	}

	wg.Wait()
	if len(repoErrs) > 0 {
		return repoErrs
	}
	return nil
}

func (c *Client) ListCharts(repoName string) ([]*vo.Chart, error) {
	entries, err := c.loadRepositoryCacheFile(repoName)
	if err != nil {
		return nil, err
	}

	var chartVOs []*vo.Chart
	for _, chart := range entries {
		latestChart := chart[chart.Len()-1]
		chartVO := &vo.Chart{
			Name:        latestChart.Name,
			Version:     latestChart.Version,
			AppVersion:  latestChart.AppVersion,
			Description: latestChart.Description,
		}
		chartVOs = append(chartVOs, chartVO)
	}
	return chartVOs, nil
}

func (c *Client) ListAppsByNamespace(kubeInfo *KubeInfo, namespace string) ([]*vo.Release, error) {
	actionConfig, err := c.actionConfigInit(kubeInfo)
	if err != nil {
		return nil, err
	}
	actions := action.NewList(actionConfig)
	actions.Deployed = true
	releases, err := actions.Run()

	var releaseVOs []*vo.Release
	for _, release := range releases {
		releaseVO := &vo.Release{
			Name:       release.Name,
			Namespace:  release.Namespace,
			Revision:   release.Version,
			Updated:    release.Info.LastDeployed.Time,
			Status:     release.Info.Status.String(),
			Chart:      release.Chart.Name(),
			AppVersion: release.Chart.AppVersion(),
		}
		releaseVOs = append(releaseVOs, releaseVO)
	}
	return releaseVOs, err
}

func (c *Client) ListAppsByAllNamespaces(kubeInfo *KubeInfo) ([]*vo.Release, error) {
	actionConfig, err := c.actionConfigInit(kubeInfo)
	if err != nil {
		return nil, err
	}
	actions := action.NewList(actionConfig)
	actions.Deployed = true
	actions.AllNamespaces = true

	releases, err := actions.Run()

	var releaseVOs []*vo.Release
	for _, release := range releases {
		releaseVO := &vo.Release{
			Name:       release.Name,
			Namespace:  release.Namespace,
			Revision:   release.Version,
			Updated:    release.Info.LastDeployed.Time,
			Status:     release.Info.Status.String(),
			Chart:      release.Chart.Name(),
			AppVersion: release.Chart.AppVersion(),
		}
		releaseVOs = append(releaseVOs, releaseVO)
	}
	return releaseVOs, err
}

func InitKubeInfo(namespace, context, config string) *KubeInfo {
	return &KubeInfo{
		namespace: namespace,
		context:   context,
		config:    config,
	}
}
