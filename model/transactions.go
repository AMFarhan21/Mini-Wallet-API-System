package model

import "time"

type Transactions struct {
	ID          int       `json:"id"`
	WalletId    int       `json:"wallet_id"`
	Type        string    `json:"type"`
	Amount      int64     `json:"amount"`
	ReferenceId string    `json:"reference_id"`
	CreatedAt   time.Time `json:"created_at"`
}
