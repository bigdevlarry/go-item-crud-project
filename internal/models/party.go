package models

type Party struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  string  `json:"last_name" binding:"required"`
	Account   Account `json:"account" binding:"required"`
}
