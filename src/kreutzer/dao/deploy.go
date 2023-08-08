package dao

type Environment struct {
	BaseDao
	Name        string `gorm:"column:name"`
	Key         string `gorm:"column:key"`
	Description string `gorm:"column:description"`
}

func (Environment) TableName() string {
	return "environment"
}
