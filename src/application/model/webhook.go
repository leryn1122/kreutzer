package model

type Webhook struct {
	BaseModel
	HookID  string `gorm:"column:hook_id"`
	Channel string `gorm:"column:channel"`
	Enabled bool   `gorm:"column:enabled"`
	URL     string `gorm:"column:url"`
	Secret  string `gorm:"column:secret"`
}

func (Webhook) TableName() string {
	return "webhook"
}
