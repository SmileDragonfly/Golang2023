package bussiness

import (
	"context"
	"sqlctest/model"
)

type IBiz interface {
	CreateWallet(ctx context.Context, wallet *model.Wallet) error
	ReadWallet(ctx context.Context, mobile string) (*model.Wallet, error)
}
