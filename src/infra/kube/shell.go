package kube

import (
	"context"
	"fmt"
	"github.com/go-errors/errors"
	ws "github.com/leryn1122/kreutzer/v2/infra/websocket"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
	"os"
)

//type StreamHandler struct {
//
//}

type ExecuteShellParam struct {
	ClusterID     string
	Namespace     string
	PodName       string
	ContainerName string
}

type StreamHandler struct {
	Stream      *ws.Connection
	ResizeEvent chan remotecommand.TerminalSize
}

func (s StreamHandler) Next() *remotecommand.TerminalSize {
	ret := <-s.ResizeEvent
	return &ret
}

func createSPDYExecutorForShell(config *ManagedKubeConfig, params ExecuteShellParam) (remotecommand.Executor, error) {
	if params.Namespace == "" {
		params.Namespace = "default"
	}
	if params.ClusterID == "" {
		params.ClusterID = "default"
	}
	if params.PodName == "" {
		return nil, errors.Errorf("pod name required")
	}
	if params.ContainerName == "" {
		return nil, errors.Errorf("container name required")
	}

	restConfig, err := BuildRestConfigFromManaged(config)
	if err != nil {
		return nil, err
	}
	client, err := NewClientFromManaged(config)
	if err != nil {
		return nil, err
	}

	request := client.CoreV1().RESTClient().
		Post().
		Resource("pods").
		Namespace(params.Namespace).
		Name(params.PodName).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: params.ContainerName,
			Command: []string{
				//				"/bin/sh",
				//				"-c",
				//				`TERM=xterm-256color; export TERM; [ -x /bin/bash ] && \
				//([ -x /usr/bin/script ] && /usr/bin/script -q -c "/bin/bash" /dev/null || exec /bin/bash) \
				//|| exec /bin/sh`,
				"/bin/bash",
			},
			Stdin:  true,
			Stdout: true,
			Stderr: true,
			TTY:    true,
		}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(restConfig, http.MethodPost, request.URL())
	return executor, nil
}

func ExecShell(config *ManagedKubeConfig, params ExecuteShellParam, conn *ws.Connection) error {
	executor, err := createSPDYExecutorForShell(config, params)

	handler := &StreamHandler{
		Stream:      conn,
		ResizeEvent: make(chan remotecommand.TerminalSize, 0),
	}

	go func() {
		err = executor.StreamWithContext(context.Background(), remotecommand.StreamOptions{
			Stdin:  handler,
			Stdout: handler,
			Stderr: os.Stderr,
			Tty:    true,
		})
		if err != nil {
			logrus.Errorf("failed to communicate with connection: %+v", err)
		}
	}()

	return err
}

// Read StreamHandler read from websocket and write into shell.
func (s StreamHandler) Read(p []byte) (size int, err error) {
	message, err := s.Stream.ReadMessage()
	if err != nil {
		logrus.Errorf("failed to read from websocket: %+v", err)
		return -1, err
	}
	result := make([]byte, len(p))
	size = len(message.Data)
	copy(p, result)
	fmt.Printf("read : %s\n", string(message.Data))
	return size, err
}

// Write StreamHandler read from shell and write into shell websocket.
func (s StreamHandler) Write(p []byte) (size int, err error) {
	data := make([]byte, len(p))
	size = len(p)
	copy(data, p)
	fmt.Printf("write : %s ==== %s\n", string(p), p)
	err = s.Stream.WriteMessage(ws.Message{
		Type: ws.BinaryMessageType,
		Data: data,
	})
	if err != nil {
		logrus.Errorf("failed to write to websocket: %+v", err)
		return -1, err
	}
	return size, nil
}

type Nop struct{}

func (n Nop) Read(p []byte) (size int, err error) {
	return 0, nil
}

func (n Nop) Write(p []byte) (size int, err error) {
	return 0, nil
}
