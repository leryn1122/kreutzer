package adapter

import (
	"github.com/leryn1122/kreutzer/v2/application/repository"
	"github.com/leryn1122/kreutzer/v2/infra/helm"
)

type helmClient struct {
	client *helm.Client
}

func NewHelmDriver() repository.HelmClient {
	return helmClient{
		client: helm.NewHelmClient(),
	}
}
