package bussiness

import (
	"context"
	"gormtest/GormTest/model"
)

type deleteStorage interface {
	deleteWallet(ctx context.Context, data *model.Wallet) error
	deleteBank(ctx context.Context, data *model.Bank) error
	deleteLinkInfo(ctx context.Context, data *model.LinkInfo) error
}

type deleteBiz struct {
	store deleteStorage
}

func (c deleteBiz) DeleteWallet(ctx context.Context, data *model.Wallet) error {
	return c.store.deleteWallet(ctx, data)
}

func (c deleteBiz) DeleteBank(ctx context.Context, data *model.Bank) error {
	return c.store.deleteBank(ctx, data)
}

func (c deleteBiz) DeleteLinkInfo(ctx context.Context, data *model.LinkInfo) error {
	return c.store.deleteLinkInfo(ctx, data)
}
