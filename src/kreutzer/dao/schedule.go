package dao

type Schedule struct {
	BaseDao
	Domain     string `json:"domain" gorm:"column:domain"`
	EntryKey   string `json:"entry_key" gorm:"column:entry_key"`
	EntryValue string `json:"entry_value" gorm:"column:entry_value"`
}

func (Schedule) TableName() string {
	return "schedule"
}
