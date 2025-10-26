package entities

type Account struct {
	SortCode      string `json:"sort_code" binding:"required,sortcode"`
	AccountNumber string `json:"account_number" binding:"required,len=8"`
}
