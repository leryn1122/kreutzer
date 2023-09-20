package tekton

import (
	"context"
	"github.com/sirupsen/logrus"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

//goland:noinspection GoNameStartsWithPackageName
const (
	TektonNamespace         = "tekton-pipelines"
	TektonResolverNamespace = "tekton-pipelines-resolvers"
)

type Client struct {
}

func (c *Client) CreateTask(task *tekton.Task) {
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Errorf("error creating Kubernetes config: %+v", err)
		return
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		logrus.Errorf("error creating Tekton client: %+v", err)
		return
	}

	createdTask, err := clientset.TektonV1().
		Tasks(TektonNamespace).
		Create(context.Background(), task, metav1.CreateOptions{})
	if err != nil {
		logrus.Errorf("failed to creating Task: %+v", err)
		return
	}

	logrus.Info("Task created successfully!")
	logrus.Info(createdTask.Name)
}
