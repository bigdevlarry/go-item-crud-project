package models

type Account struct {
	SortCode      string `json:"sort_code" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
}
