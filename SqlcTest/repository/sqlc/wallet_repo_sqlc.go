package sqlc

import (
	"context"
	"sqlctest/model"
	sqlc_postgresql "sqlctest/repository/sqlc/postgresql/db"
)

type WalletRepoSqlc struct {
	db sqlc_postgresql.DBTX
}

func NewWalletRepoSqlc(db sqlc_postgresql.DBTX) WalletRepoSqlc {
	return WalletRepoSqlc{db}
}

func (r WalletRepoSqlc) CreateWallet(ctx context.Context, wallet *model.Wallet) error {
	query := sqlc_postgresql.New(r.db)
	return query.InsertWallet(ctx, sqlc_postgresql.InsertWalletParams{wallet.UserName, wallet.Password, wallet.Mobile})
}

func (r WalletRepoSqlc) ReadWallet(ctx context.Context, mobile string) (*model.Wallet, error) {
	query := sqlc_postgresql.New(r.db)
	wallet, err := query.ReadWallet(ctx, mobile)
	if err != nil {
		return nil, err
	}
	var walletRet model.Wallet
	walletRet.ID = wallet.ID
	wallet.Createtime = wallet.Createtime
	walletRet.UserName = wallet.Username
	walletRet.Mobile = wallet.Mobile
	return &walletRet, nil
}
