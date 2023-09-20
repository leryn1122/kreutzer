package kube

import (
  kube "github.com/leryn1122/kreutzer/v2/infra/kubeconfig"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeInfo struct {
	Cluster   string
	Namespace string
}

// ManagedKubeConfig
// A kube config managed by application
type ManagedKubeConfig struct {
	ID       string
	URL      string
	cert     string
	username string
	token    string
}

func NewKubeInfo(cluster, namespace string) *KubeInfo {
	return &KubeInfo{
		Cluster:   cluster,
		Namespace: namespace,
	}
}

func NewManagedKubeConfig(id, url, cert, username, token string) *ManagedKubeConfig {
	return &ManagedKubeConfig{
		ID:       id,
		URL:      url,
		cert:     cert,
		username: username,
		token:    token,
	}
}

func BuildRestConfigFromManaged(config *ManagedKubeConfig) (*rest.Config, error) {
	kubeConfigText, err := kube.ForTokenBased(config.ID, config.URL, config.cert, config.username, config.token)
	if err != nil {
		return nil, err
	}
	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfigText))
	return kubeConfig, nil
}

func NewClientFromManaged(config *ManagedKubeConfig) (*kubernetes.Clientset, error) {
	kubeConfigText, err := kube.ForTokenBased(config.ID, config.URL, config.cert, config.username, config.token)
	if err != nil {
		return nil, err
	}
	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfigText))
	if err != nil {
		logrus.Errorf("failed to create the kube config: %+v", err)
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		logrus.Errorf("failed to create kube client: %+v", err)
		return nil, err
	}
	return clientset, err
}

// NewClientOrDie Create a client or panic.
func NewClientOrDie(config *ManagedKubeConfig) *kubernetes.Clientset {
	kubeConfigText, err := kube.ForTokenBased(config.ID, config.URL, config.cert, config.username, config.token)
	if err != nil {
		panic(err)
	}
	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfigText))
	if err != nil {
		logrus.Errorf("failed to create the kube config: %+v", err)
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		logrus.Errorf("failed to create kube client: %+v", err)
		panic(err)
	}
	return clientset
}
