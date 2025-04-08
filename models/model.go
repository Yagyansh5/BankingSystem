package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID          int64              `json:"id" gorm:"primaryKey"`
	Owner       string             `json:"owner"`
	Balance     int64              `json:"balance"`
	Currency    string             `json:"currency"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	CountryCode pgtype.Int4        `json:"country_code"`
}

type Entry struct {
	ID        int64              `json:"id" gorm:"primaryKey"`
	AccountID int64              `json:"account_id"`
	Amount    int64              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Transfer struct {
	ID            int64              `json:"id" gorm:"primaryKey"`
	FromAccountID int64              `json:"from_account_id"`
	ToAccountID   int64              `json:"to_account_id"`
	Amount        int64              `json:"amount"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}
