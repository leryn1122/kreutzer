package dao

import "time"

type HelmRepo struct {
	BaseDao
	Name                  string    `gorm:"column:name"`
	URL                   string    `gorm:"column:url"`
	LastSyncTime          time.Time `gorm:"column:last_sync_time"`
	Username              string    `gorm:"column:username"`
	Password              string    `gorm:"column:password"`
	PassCredentialsAll    bool      `gorm:"column:pass_credentials_all"`
	InsecureSkipTLSVerify bool      `gorm:"column:insecure_skip_tls_verify"`
	CertFile              string    `gorm:"column:cert_file"`
	KeyFile               string    `gorm:"column:key_file"`
	CAFile                string    `gorm:"column:ca_file"`
}

func (HelmRepo) TableName() string {
	return "helm_repo"
}
