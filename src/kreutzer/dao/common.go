package dao

import "time"

type BaseDao struct {
	Id           int       `gorm:"column:id;unique;primaryKey;autoIncrement"`
	CreationTime time.Time `gorm:"column:creation_time"`
	ModifiedTime time.Time `gorm:"column:modified_time"`
}
