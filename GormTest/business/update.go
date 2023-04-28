package business

import (
	"context"
	"gormtest/GormTest/model"
	"gormtest/GormTest/storage"
)

type updateBiz struct {
	store storage.Storage
}

func NewUpdateBiz(store storage.Storage) updateBiz {
	return updateBiz{store: store}
}

func (u updateBiz) UpdateWallet(ctx context.Context, condition map[string]interface{}, dataUpdate *model.Wallet) error {
	return u.store.UpdateWallet(ctx, condition, dataUpdate)
}

func (u updateBiz) UpdateBank(ctx context.Context, condition map[string]interface{}, dataUpdate *model.Bank) error {
	return u.store.UpdateBank(ctx, condition, dataUpdate)
}

func (u updateBiz) UpdateLinkInfo(ctx context.Context, condition map[string]interface{}, dataUpdate *model.LinkInfo) error {
	return u.store.UpdateLinkInfo(ctx, condition, dataUpdate)
}
