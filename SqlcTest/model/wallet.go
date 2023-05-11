package model

import "time"

type Wallet struct {
	ID         int64
	CreateTime time.Time
	UserName   string
	Password   string
	Mobile     string
}
