package dao

type Environment struct {
	BaseDao
	Name        string `json:"name" gorm:"column:name"`
	Key         string `json:"key" gorm:"column:key"`
	Description string `json:"description" gorm:"column:description"`
}

func (Environment) TableName() string {
	return "environment"
}
