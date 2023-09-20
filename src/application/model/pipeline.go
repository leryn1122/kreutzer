package model

type Pipeline struct {
	BaseModel
	PipelineId string `gorm:"column:pipeline_id"`
	Name       string `gorm:"column:name"`
	Repo       string `gorm:"column:repo"`
	Enabled    bool   `gorm:"column:enabled"`
}

func (Pipeline) TableName() string {
	return "pipeline"
}
