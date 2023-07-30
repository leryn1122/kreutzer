package dao

import "time"

type BaseDao struct {
	Id           int       `json:"id" gorm:"column:id;unique;primaryKey;autoIncrement"`
	CreationTime time.Time `json:"creation_time" gorm:"column:creation_time"`
	ModifiedTime time.Time `json:"modified_time" gorm:"column:modified_time"`
}
