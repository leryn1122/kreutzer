package kubeApi

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/infra/kube"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPersistenceVolumes(config *kube.ManagedKubeConfig) (*v1.PersistentVolumeList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	persistentVolumes, err := client.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return persistentVolumes, nil
}

func ListPersistenceVolumeClaims(config *kube.ManagedKubeConfig, namespace string) (*v1.PersistentVolumeClaimList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	persistentVolumeClaims, err := client.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return persistentVolumeClaims, nil
}
