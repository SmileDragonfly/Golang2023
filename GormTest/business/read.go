package business

import (
	"context"
	"gormtest/GormTest/model"
	"gormtest/GormTest/storage"
)

type readBiz struct {
	store storage.Storage
}

func NewReadBiz(store storage.Storage) readBiz {
	return readBiz{store: store}
}

func (r readBiz) ReadWallet(ctx context.Context, condition map[string]interface{}) ([]model.Wallet, error) {
	return r.store.ReadWallet(ctx, condition)
}

func (r readBiz) ReadBank(ctx context.Context, condition map[string]interface{}) ([]model.Bank, error) {
	return r.store.ReadBank(ctx, condition)
}

func (r readBiz) ReadLinkInfo(ctx context.Context, condition map[string]interface{}) ([]model.LinkInfo, error) {
	return r.store.ReadLinkInfo(ctx, condition)
}
