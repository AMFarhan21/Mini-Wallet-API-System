package model

type Wallets struct {
	ID      int   `json:"id"`
	UserId  int   `json:"user_id"`
	Balance int64 `json:"balance"`
}
