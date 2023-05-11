package bussiness

import (
	"context"
	"sqlctest/model"
	"sqlctest/repository"
)

type Biz struct {
	repo repository.IWallet
}

func (r Biz) CreateWallet(ctx context.Context, wallet *model.Wallet) error {
	return r.repo.CreateWallet(ctx, wallet)
}

func (r Biz) ReadWallet(ctx context.Context, mobile string) (*model.Wallet, error) {
	return r.repo.ReadWallet(ctx, mobile)
}
