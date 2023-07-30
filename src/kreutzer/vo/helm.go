package vo

import "time"

type Repo struct {
	Name         string    `json:"name"`
	URL          string    `json:"URL"`
	LastSyncTime time.Time `json:"lastSyncTime"`
}

type Chart struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	AppVersion  string `json:"appVersion"`
	Description string `json:"description"`
}

type Release struct {
	Name       string    `json:"name"`
	Namespace  string    `json:"namespace"`
	Revision   int       `json:"revision"`
	Updated    time.Time `json:"updated"`
	Status     string    `json:"status"`
	Chart      string    `json:"chart"`
	AppVersion string    `json:"appVersion"`
}
