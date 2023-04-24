package model

import "time"

type Wallet struct {
	ID           int       `gorm:"autoIncrement;primaryKey"`
	CreateTime   time.Time `gorm:"autoCreateTime:milli"`
	UpdateTime   time.Time `gorm:"autoUpdateTime:milli"`
	Mobile       string    `gorm:"size:16"`
	Email        string    `gorm:"size:255"`
	NationalID   string    `gorm:"size:32"`
	UserName     string    `gorm:"size:32;unique"`
	Password     string    `gorm:"size:512"`
	PasswordHash string    `gorm:"size:32"`
	PIN          string    `gorm:"size:16"`
	Fullname     string    `gorm:"size:64"`
	Address      string    `gorm:"size:512"`
	Balance      int64
	Language     string `gorm:"size:16"`
	Gender       string `gorm:"size:8"`
	AuthenType   string `gorm:"size:8"`
}
