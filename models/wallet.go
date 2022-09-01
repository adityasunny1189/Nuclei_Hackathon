package models

type Wallet struct {
	Id      int     `json:"id"`
	UserId  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}
