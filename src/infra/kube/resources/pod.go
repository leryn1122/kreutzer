package kubeApi

import (
	"context"
	"github.com/leryn1122/kreutzer/v2/infra/kube"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPods(config *kube.ManagedKubeConfig, namespace string) (*v1.PodList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

func ListDeployments(config *kube.ManagedKubeConfig, namespace string) (*appv1.DeploymentList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	deploys, err := client.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return deploys, nil
}

func ListStatefulSets(config *kube.ManagedKubeConfig, namespace string) (*appv1.StatefulSetList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	statefulSets, err := client.AppsV1().StatefulSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return statefulSets, nil
}

func ListDaemonSets(config *kube.ManagedKubeConfig, namespace string) (*appv1.DaemonSetList, error) {
	client, err := kube.NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}
	daemonSets, err := client.AppsV1().DaemonSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return daemonSets, nil
}
