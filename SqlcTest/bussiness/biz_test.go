package bussiness

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"reflect"
	"sqlctest/model"
	"sqlctest/repository"
	"sqlctest/repository/sqlc"
	"testing"
)

func TestBiz_CreateWallet(t *testing.T) {
	type fields struct {
		repo repository.IWallet
	}
	type args struct {
		ctx    context.Context
		wallet *model.Wallet
	}
	strConn := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		5432, "127.0.0.1", "postgres", "123@123A", "sqlctest")
	strDriver := "postgres"
	db, err := sql.Open(strDriver, strConn)
	if err != nil {
		t.Errorf("Connect to db error:%s", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"CreateWallet1",
			fields{sqlc.NewWalletRepoSqlc(db)},
			args{context.Background(), &model.Wallet{UserName: "DatDT", Password: "123456", Mobile: "0312658794"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Biz{
				repo: tt.fields.repo,
			}
			if err := r.CreateWallet(tt.args.ctx, tt.args.wallet); (err != nil) != tt.wantErr {
				t.Errorf("CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBiz_ReadWallet(t *testing.T) {
	type fields struct {
		repo repository.IWallet
	}
	type args struct {
		ctx    context.Context
		mobile string
	}
	strConn := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		5432, "127.0.0.1", "postgres", "123@123A", "sqlctest")
	strDriver := "postgres"
	db, err := sql.Open(strDriver, strConn)
	if err != nil {
		t.Errorf("Connect to db error:%s", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"TestReadWallet1",
			fields{sqlc.NewWalletRepoSqlc(db)},
			args{context.Background(), "0312658794"},
			"DatDT",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Biz{
				repo: tt.fields.repo,
			}
			got, err := r.ReadWallet(tt.args.ctx, tt.args.mobile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadWallet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.UserName, tt.want) {
				t.Errorf("ReadWallet() got = %v, want %v", got.UserName, tt.want)
			}
		})
	}
}
