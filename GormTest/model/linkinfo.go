package model

import "time"

type LinkInfo struct {
	ID          int       `gorm:"autoIncrement;primaryKey"`
	CreateTime  time.Time `gorm:"autoCreateTime:milli"`
	UpdateTime  time.Time `gorm:"autoUpdateTime:milli"`
	WalletID    int
	Wallet      Wallet
	BankID      int
	Bank        Bank
	BankAccount string `gorm:"size:32"`
	BankCard    string `gorm:"size:32"`
	ResourceID  string `gorm:"size:36"`
}
