package business

import (
	"context"
	"gormtest/GormTest/model"
)

type createStorage interface {
	createWallet(ctx context.Context, data *model.Wallet) error
	createBank(ctx context.Context, data *model.Bank) error
	createLinkInfo(ctx context.Context, data *model.LinkInfo) error
}

type createBiz struct {
	store createStorage
}

func (c createBiz) CreateWallet(ctx context.Context, data *model.Wallet) error {
	return c.store.createWallet(ctx, data)
}

func (c createBiz) CreateBank(ctx context.Context, data *model.Bank) error {
	return c.store.createBank(ctx, data)
}

func (c createBiz) CreateLinkInfo(ctx context.Context, data *model.LinkInfo) error {
	return c.store.createLinkInfo(ctx, data)
}
