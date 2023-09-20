package model

type Environment struct {
	BaseModel
	Name        string `gorm:"column:name"`
	Key         string `gorm:"column:key"`
	Description string `gorm:"column:description"`
}

func (Environment) TableName() string {
	return "environment"
}
