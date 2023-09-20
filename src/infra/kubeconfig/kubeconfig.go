package kube

import (
	"bytes"
	"html/template"
)

type kubeConfigTemplate struct {
	ClusterID string
	URL       string
	Cert      string
	User      string
	Token     string
}

func ForTokenBased(clusterID, url, cert, username, token string) (string, error) {
	data := &kubeConfigTemplate{
		ClusterID: clusterID,
		URL:       url,
		Cert:      cert,
		User:      username,
		Token:     token,
	}

	buf := &bytes.Buffer{}
	err := tokenTemplate.Execute(buf, data)
	return buf.String(), err
}

const (
	tokenTemplateText = `
apiVersion: v1
kind: Config
clusters:
- name: "{{ .ClusterID }}"
  cluster:
    server: "{{ .URL }}"
    certificate-authority-data: {{ .Cert }}
contexts:
- name: "{{ .User }}@{{ .ClusterID }}"
  context:
    cluster: "{{ .ClusterID }}"
    user: "{{ .User }}"
    namespace: "{{ .User }}"
current-context: "{{ .User }}@{{ .ClusterID }}"
users:
- name: "{{ .User }}"
  user:
    token: {{ .Token }}`
)

var (
	tokenTemplate = template.Must(template.New("tokenTemplate").Parse(tokenTemplateText))
)
