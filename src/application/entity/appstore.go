package entity

import (
	"github.com/leryn1122/kreutzer/v2/application/vo"
	"github.com/leryn1122/kreutzer/v2/infra/id"
	"time"
)

type AppStore struct {
	ID       id.Descriptor `json:"id"`
	Name     string        `json:"name"`
	URL      string        `json:"url"`
	LastSync time.Time     `json:"lastSync"`
}

func (e *AppStore) ToViewFromAppStore() *vo.AppStore {
	return &vo.AppStore{
		Name:     e.Name,
		URL:      e.URL,
		LastSync: e.LastSync,
	}
}

type AppRelease struct {
	Name       string    `json:"name"`
	Namespace  string    `json:"namespace"`
	Revision   int       `json:"revision"`
	Updated    time.Time `json:"updated"`
	Status     string    `json:"status"`
	Chart      string    `json:"chart"`
	AppVersion string    `json:"appVersion"`
}
