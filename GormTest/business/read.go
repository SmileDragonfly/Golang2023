package business

import (
	"context"
	"gormtest/GormTest/model"
)

type readStorage interface {
	readWallet(ctx context.Context, data *model.Wallet) error
	readBank(ctx context.Context, data *model.Bank) error
	readLinkInfo(ctx context.Context, data *model.LinkInfo) error
}

type readBiz struct {
	store readStorage
}

func (c readBiz) ReadWallet(ctx context.Context, data *model.Wallet) error {
	return c.store.readWallet(ctx, data)
}

func (c readBiz) readBank(ctx context.Context, data *model.Bank) error {
	return c.store.readBank(ctx, data)
}

func (c readBiz) readLinkInfo(ctx context.Context, data *model.LinkInfo) error {
	return c.store.readLinkInfo(ctx, data)
}
