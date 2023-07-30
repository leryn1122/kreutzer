package tekton

//import (
//	"context"
//	"fmt"
//	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
//	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/rest"
//)
//
//const (
//	TektonNamespace         = "tekton-pipelines"
//	TektonResolverNamespace = "tekton-pipelines-resolvers"
//)
//
//type Client struct {
//}
//
//func CreateTask(task *v1.Task) {
//	config, err := rest.InClusterConfig()
//	if err != nil {
//		fmt.Println("Error creating Kubernetes config:", err)
//		return
//	}
//
//	clientset, err := versioned.NewForConfig(config)
//	if err != nil {
//		fmt.Println("Error creating Tekton client:", err)
//		return
//	}
//
//	createdTask, err := clientset.TektonV1().Tasks(TektonNamespace).
//		Create(context.Background(), task, metav1.CreateOptions{})
//	if err != nil {
//		fmt.Println("Error creating Task:", err)
//		return
//	}
//
//	fmt.Println("Task created successfully!")
//	fmt.Println(createdTask.Name)
//}
//
//func GenerateTask(name string) *v1.Task {
//	return &v1.Task{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      name,
//			Namespace: TektonNamespace,
//		},
//		Spec: v1.TaskSpec{
//			Steps: []v1.Step{
//				{
//					Name:    "",
//					Image:   "",
//					Command: []string{},
//				},
//			},
//		},
//	}
//}
