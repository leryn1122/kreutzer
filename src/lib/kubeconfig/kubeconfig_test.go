package kube

import (
	"fmt"
	"testing"
)

func TestForTokenBased(t *testing.T) {
	kubeConfig, err := ForTokenBased("default", "default", "localhost", "default", "default")
	fmt.Println(kubeConfig)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestForBasic(t *testing.T) {
	kubeConfig, err := ForBasic("localhost", "admin", "admin@123")
	fmt.Println(kubeConfig)
	if err != nil {
		t.Errorf(err.Error())
	}
}
