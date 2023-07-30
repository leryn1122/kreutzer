package dao

type Pipeline struct {
	BaseDao
	PipelineId string `json:"pipeline_id" gorm:"column:pipeline_id"`
	Name       string `json:"name" gorm:"column:name"`
	Repo       string `json:"repo" gorm:"column:repo"`
	Enabled    bool   `json:"enabled" gorm:"column:enabled"`
}

func (Pipeline) TableName() string {
	return "pipeline"
}
