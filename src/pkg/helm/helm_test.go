package helm

import (
	"os"
	"testing"
)

func TestForHelmListApp(t *testing.T) {
	kubeInfo := InitKubeInfo("kube-system", "default", os.ExpandEnv("$HOME/.kube/config"))
	client := NewHelmClient()
	releases, err := client.ListAppsByNamespace(kubeInfo, "kube-system")
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, release := range releases {
		print(release)
	}
}
