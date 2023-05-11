-- name: InsertWallet :exec
INSERT INTO Wallet(UserName, Password, Mobile) VALUES ($1, $2, $3);
-- name: ReadWallet :one
SELECT * FROM Wallet WHERE Mobile=$1;