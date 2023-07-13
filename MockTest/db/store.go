package db

import "mocktest/model"

type Store interface {
	CreateUser(mobile string, password string) (*model.User, error)
	CreateMerchant(name string, code string) (*model.Merchant, error)
}
