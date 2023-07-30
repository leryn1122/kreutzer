package dao

import "time"

type HelmRepo struct {
	BaseDao
	Name                  string    `gorm:"column:name"`
	URL                   string    `gorm:"column:url"`
	LastSyncTime          time.Time `gorm:"column:last_sync_time"`
	Username              string    `gorm:"username"`
	Password              string    `gorm:"password"`
	PassCredentialsAll    bool      `gorm:"pass_credentials_all"`
	InsecureSkipTLSVerify bool      `gorm:"insecure_skip_tls_verify"`
	CertFile              string    `gorm:"cert_file"`
	KeyFile               string    `gorm:"key_file"`
	CAFile                string    `gorm:"ca_file"`
}

func (HelmRepo) TableName() string {
	return "helm_repo"
}
