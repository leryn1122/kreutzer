package dao

type Schedule struct {
	BaseDao
	CronExpression string `gorm:"column:cron_expression"`
}

func (Schedule) TableName() string {
	return "schedule"
}
