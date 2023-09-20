package model

type FlexCube struct {
	BaseModel
	Domain     string `gorm:"column:domain"`
	EntryKey   string `gorm:"column:entry_key"`
	EntryValue string `gorm:"column:entry_value"`
}

func (FlexCube) TableName() string {
	return "flex_cube"
}
