package dao

type Webhook struct {
	BaseDao
	WebhookId string `gorm:"column:webhook_id"`
	Channel   string `gorm:"column:channel"`
	Enabled   bool   `gorm:"column:enabled"`
	URL       string `gorm:"column:url"`
	Secret    string `gorm:"column:secret"`
}

func (Webhook) TableName() string {
	return "webhook"
}
