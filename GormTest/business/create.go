package business

import (
	"context"
	"gormtest/GormTest/model"
	"gormtest/GormTest/storage"
)

type createBiz struct {
	store storage.Storage
}

func NewCreateBiz(store storage.Storage) createBiz {
	return createBiz{store: store}
}

func (c createBiz) CreateWallet(ctx context.Context, data *model.Wallet) error {
	return c.store.CreateWallet(ctx, data)
}

func (c createBiz) CreateBank(ctx context.Context, data *model.Bank) error {
	return c.store.CreateBank(ctx, data)
}

func (c createBiz) CreateLinkInfo(ctx context.Context, data *model.LinkInfo) error {
	return c.store.CreateLinkInfo(ctx, data)
}
