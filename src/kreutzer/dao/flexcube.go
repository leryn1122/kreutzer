package dao

type FlexCube struct {
	BaseDao
	Domain     string `json:"domain" gorm:"column:domain"`
	EntryKey   string `json:"entry_key" gorm:"column:entry_key"`
	EntryValue string `json:"entry_value" gorm:"column:entry_value"`
}

func (FlexCube) TableName() string {
	return "flex_cube"
}
