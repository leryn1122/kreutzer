package kubeApi

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/infra/kube"
	v1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListServices(config *kube.ManagedKubeConfig, namespace string) (*v1.ServiceList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	services, err := client.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return services, nil
}

func ListIngresses(config *kube.ManagedKubeConfig, namespace string) (*networkv1.IngressList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	ingresses, err := client.NetworkingV1().Ingresses(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return ingresses, nil
}
