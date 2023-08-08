package dao

type ManagedCluster struct {
	BaseDao
	Name     string `gorm:"column:name"`
	URL      string `gorm:"column:url"`
	Enabled  string `gorm:"column:enabled"`
	Cert     string `gorm:"column:cert"`
	Username string `gorm:"column:username"`
	Token    string `gorm:"column:token"`
}

func (ManagedCluster) TableName() string {
	return "managed_cluster"
}
