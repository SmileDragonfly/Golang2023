package business

import (
	"context"
	"gormtest/GormTest/model"
	"gormtest/GormTest/storage"
)

type Bussiness struct {
	CreateBussiness
	DeleteBussiness
	UpdateBussiness
	ReadBussiness
	store storage.Storage
}

type CreateBussiness interface {
	CreateWallet(ctx context.Context, data *model.Wallet) error
	CreateBank(ctx context.Context, data *model.Bank) error
	CreateLinkInfo(ctx context.Context, data *model.LinkInfo) error
}

type DeleteBussiness interface {
	DeleteWallet(ctx context.Context, condition map[string]interface{}) error
	DeleteBank(ctx context.Context, condition map[string]interface{}) error
	DeleteLinkInfo(ctx context.Context, condition map[string]interface{}) error
}

type UpdateBussiness interface {
	UpdateWallet(ctx context.Context, condition map[string]interface{}, dataUpdate *model.Wallet) error
	UpdateBank(ctx context.Context, condition map[string]interface{}, dataUpdate *model.Bank) error
	UpdateLinkInfo(ctx context.Context, condition map[string]interface{}, dataUpdate *model.LinkInfo) error
}

type ReadBussiness interface {
	ReadWallet(ctx context.Context, condition map[string]interface{}) ([]model.Wallet, error)
	ReadBank(ctx context.Context, condition map[string]interface{}) ([]model.Bank, error)
	ReadLinkInfo(ctx context.Context, condition map[string]interface{}) ([]model.LinkInfo, error)
}
