package helm

import (
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
	"os"
	"path/filepath"
	"sync"
)

type RepoErr struct {
	Name string
	Err  string
}

func LoadCachedRepos() (*repo.File, error) {
	path := cli.New().EnvVars()["HELM_REPOSITORY_CONFIG"]
	return repo.LoadFile(os.ExpandEnv(path))
}

func LoadCachedRepoIndex(name string) (*repo.IndexFile, error) {
	path := filepath.Join(
		cli.New().EnvVars()["HELM_REPOSITORY_CACHE"],
		name+"-index.yaml",
	)
	path = os.ExpandEnv(path)
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return repo.LoadIndexFile(os.ExpandEnv(path))
}

func UpdateRepo(entry *repo.Entry) error {
	repo_, err := repo.NewChartRepository(entry, getter.All(cli.New()))
	if err != nil {
		return err
	}
	_, err = repo_.DownloadIndexFile()
	if err != nil {
		return err
	}
	return nil
}

func UpdateRepos() []RepoErr {
	var repoErrs []RepoErr
	file, err := LoadCachedRepos()
	if err != nil {
	}
	var wg sync.WaitGroup
	for _, repo_ := range file.Repositories {
		wg.Add(1)
		go func(entry *repo.Entry) {
			defer wg.Done()
			err := UpdateRepo(entry)
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
