package kubeApi

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/lib/kube"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListConfigMaps(config *kube.ManagedKubeConfig, namespace string) (*v1.ConfigMapList, error) {
	client, err := kube.NewClient(config)
	if err != nil {
		return nil, err
	}
	configMaps, err := client.CoreV1().ConfigMaps(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return configMaps, nil
}

func ListSecrets(config *kube.ManagedKubeConfig, namespace string) (*v1.SecretList, error) {
	client, err := kube.NewClient(config)
	if err != nil {
		return nil, err
	}
	secrets, err := client.CoreV1().Secrets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return secrets, nil
}
