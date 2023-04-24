package model

import "time"

type Bank struct {
	ID         int       `gorm:"autoIncrement;primaryKey"`
	CreateTime time.Time `gorm:"autoCreateTime:milli"`
	UpdateTime time.Time `gorm:"autoUpdateTime:milli"`
	Code       string    `gorm:"size:8;unique"`
	Name       string    `gorm:"size:128"`
}
