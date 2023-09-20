package kubeApi

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/infra/kube"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNamespaces(config *kube.ManagedKubeConfig) (*v1.NamespaceList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces, nil
}
