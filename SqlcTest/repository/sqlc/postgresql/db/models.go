// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc_postgresql

import (
	"database/sql"
)

type Wallet struct {
	ID         int64
	Createtime sql.NullTime
	Username   string
	Password   string
	Mobile     string
}
