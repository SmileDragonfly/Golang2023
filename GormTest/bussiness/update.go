package bussiness

import (
	"context"
	"gormtest/GormTest/model"
)

type updateStorage interface {
	updateWallet(ctx context.Context, data *model.Wallet) error
	updateBank(ctx context.Context, data *model.Bank) error
	updateLinkInfo(ctx context.Context, data *model.LinkInfo) error
}

type updateBiz struct {
	store updateStorage
}

func (c updateBiz) UpdateWallet(ctx context.Context, data *model.Wallet) error {
	return c.store.updateWallet(ctx, data)
}

func (c updateBiz) UpdateBank(ctx context.Context, data *model.Bank) error {
	return c.store.updateBank(ctx, data)
}

func (c updateBiz) UpdateLinkInfo(ctx context.Context, data *model.LinkInfo) error {
	return c.store.updateLinkInfo(ctx, data)
}
