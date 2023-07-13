// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package sqlc_postgresql

import (
	"context"
)

const insertWallet = `-- name: InsertWallet :exec
INSERT INTO Wallet(UserName, Password, Mobile) VALUES ($1, $2, $3)
`

type InsertWalletParams struct {
	Username string
	Password string
	Mobile   string
}

func (q *Queries) InsertWallet(ctx context.Context, arg InsertWalletParams) error {
	_, err := q.db.ExecContext(ctx, insertWallet, arg.Username, arg.Password, arg.Mobile)
	return err
}

const readWallet = `-- name: ReadWallet :one
SELECT id, createtime, username, password, mobile FROM Wallet WHERE Mobile=$1
`

func (q *Queries) ReadWallet(ctx context.Context, mobile string) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, readWallet, mobile)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.Createtime,
		&i.Username,
		&i.Password,
		&i.Mobile,
	)
	return i, err
}