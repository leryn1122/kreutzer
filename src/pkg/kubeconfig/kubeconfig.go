package kube

import (
	"bytes"
	"html/template"
)

type data struct {
	ClusterName string
	ClusterID   string
	Host        string
	Cert        string
	User        string
	Username    string
	Password    string
	Token       string
}

func ForTokenBased(clusterName, clusterID, host, username, token string) (string, error) {
	data := &data{
		ClusterName: clusterName,
		ClusterID:   clusterID,
		Host:        host,
		Cert:        "",
		User:        username,
		Token:       token,
	}

	buf := &bytes.Buffer{}
	err := tokenTemplate.Execute(buf, data)
	return buf.String(), err
}

func ForBasic(host, username, password string) (string, error) {
	data := &data{
		ClusterName: "",
		Host:        host,
		Cert:        caCertString(),
		User:        username,
		Username:    username,
		Password:    password,
	}

	buf := &bytes.Buffer{}
	err := basicTemplate.Execute(buf, data)
	return buf.String(), err
}

func caCertString() string {
	return ""
}

const (
	tokenTemplateText = `
apiVersion: v1
kind: Config
clusters:
- name: {{ .ClusterName }}
  cluster:
	api-version: v1
    certificate-authority-data: {{ .Cert }}
    server: "https://{{ .Host }}/k8s/clusters/{{ .ClusterID }}"
contexts:
- name: "{{ .ClusterName }}"
  context:
    cluster: "{{ .ClusterName }}"
    user: "{{ .User }}"
current-context: "{{.ClusterName }}"
users:
- name: "{{ .User }}"
  user:
	token: {{ .Token }}
    # client-certificate-data: {{ .Cert }}
    # client-key-data: {{ .Cert }}
`

	basicTemplateText = `
apiVersion: v1
kind: Config
clusters:
- name: "{{ .ClusterName }}"
  cluster:
    api-version: v1
    server: "https://{{.Host}}"
contexts:
- name: "{{ .ClusterName }}"
  context:
    user: "{{ .User }}"
    cluster: "{{ .ClusterName }}"
current-context: "{{ .ClusterName }}"
users:
- name: "{{ .User }}"
  user:
    username: "{{ .Username }}"
    password: "{{ .Password }}"
`
)

var (
	tokenTemplate = template.Must(template.New("tokenTemplate").Parse(tokenTemplateText))
	basicTemplate = template.Must(template.New("basicTemplate").Parse(basicTemplateText))
)
