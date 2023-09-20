package model

import "time"

type AppStore struct {
	BaseModel
	Name                  string    `gorm:"column:name"`
	URL                   string    `gorm:"column:url"`
	LastSync              time.Time `gorm:"column:last_sync"`
	Username              string    `gorm:"column:username"`
	Password              string    `gorm:"column:password"`
	PassCredentialsAll    bool      `gorm:"column:pass_credentials_all"`
	InsecureSkipTLSVerify bool      `gorm:"column:insecure_skip_tls_verify"`
	CertFile              string    `gorm:"column:cert_file"`
	KeyFile               string    `gorm:"column:key_file"`
	CAFile                string    `gorm:"column:ca_file"`
}

func (AppStore) TableName() string {
	return "appstore"
}
