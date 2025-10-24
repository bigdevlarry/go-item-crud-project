package models

import "go-test/internal/models/enums"

type ItemCreateDTO struct {
	Amount     float64          `json:"amount" binding:"required,gt=0"`
	Type       enums.ItemType   `json:"type" binding:"required,itemtype"`
	Status     enums.ItemStatus `json:"status" binding:"required,itemstatus"`
	Attributes Attributes       `json:"attributes" binding:"required"`
}
