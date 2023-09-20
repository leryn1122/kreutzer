package adapter

import (
	"github.com/leryn1122/kreutzer/v2/application/repository"
)

type kubeClient struct {
}

func NewKubeClient() repository.KubeClient {
	return kubeClient{}
}
