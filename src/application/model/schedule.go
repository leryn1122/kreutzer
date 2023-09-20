package model

type Schedule struct {
	BaseModel
	CronExpression string `gorm:"column:cron_expression"`
}

func (Schedule) TableName() string {
	return "schedule"
}
