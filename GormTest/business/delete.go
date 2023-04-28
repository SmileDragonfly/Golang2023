package business

import (
	"context"
	"gormtest/GormTest/storage"
)

type deleteBiz struct {
	store storage.Storage
}

func NewDeleteBiz(store storage.Storage) deleteBiz {
	return deleteBiz{store: store}
}

func (d deleteBiz) DeleteWallet(ctx context.Context, condition map[string]interface{}) error {
	return d.store.DeleteWallet(ctx, condition)
}

func (d deleteBiz) DeleteBank(ctx context.Context, condition map[string]interface{}) error {
	return d.store.DeleteBank(ctx, condition)
}

func (d deleteBiz) DeleteLinkInfo(ctx context.Context, condition map[string]interface{}) error {
	return d.store.DeleteLinkInfo(ctx, condition)
}
