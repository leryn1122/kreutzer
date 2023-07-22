package helm

import (
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

type Client struct {
	settings   *cli.EnvSettings
	helmDriver string
}

type ClientOption = func(client *Client)

type KubeInfo struct {
	namespace string
	context   string
	config    string
}

func NewHelmClient(opts ...ClientOption) *Client {
	settings := cli.New()
	var client = &Client{
		settings:   settings,
		helmDriver: os.Getenv("HELM_DRIVER"),
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
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

func (c *Client) ListRepos(kubeInfo *KubeInfo) ([]*repo.ChartRepository, error) {
	//actionConfig, err := c.actionConfigInit(kubeInfo)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (c *Client) ListAppsByNamespace(kubeInfo *KubeInfo, namespace string) ([]*release.Release, error) {
	actionConfig, err := c.actionConfigInit(kubeInfo)
	if err != nil {
		return nil, err
	}
	actions := action.NewList(actionConfig)
	actions.Deployed = true
	return actions.Run()
}

func (c *Client) ListAppsByAllNamespaces(kubeInfo *KubeInfo) ([]*release.Release, error) {
	actionConfig, err := c.actionConfigInit(kubeInfo)
	if err != nil {
		return nil, err
	}
	actions := action.NewList(actionConfig)
	actions.Deployed = true
	actions.AllNamespaces = true
	return nil, nil
}

func InitKubeInfo(namespace, context, config string) *KubeInfo {
	return &KubeInfo{
		namespace: namespace,
		context:   context,
		config:    config,
	}
}
