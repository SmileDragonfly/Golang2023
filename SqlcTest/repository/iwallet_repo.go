package repository

import (
	"context"
	"sqlctest/model"
)

type IWallet interface {
	CreateWallet(ctx context.Context, wallet *model.Wallet) error
	ReadWallet(ctx context.Context, mobile string) (*model.Wallet, error)
}
