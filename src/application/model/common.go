package model

import "time"

type BaseModel struct {
	ID           uint64    `gorm:"column:id;unique;primaryKey;autoIncrement"`
	CreationTime time.Time `gorm:"column:creation_time"`
	ModifiedTime time.Time `gorm:"column:modified_time"`
}
