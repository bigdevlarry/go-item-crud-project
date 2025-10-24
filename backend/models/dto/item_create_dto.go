package dto

import (
	"go-test/backend/models/entities"
	"go-test/backend/models/enums"
)

type ItemCreateDTO struct {
	Amount     float64             `json:"amount" binding:"required,gt=0"`
	Type       enums.ItemType      `json:"type" binding:"required,itemtype"`
	Status     enums.ItemStatus    `json:"status" binding:"required,itemstatus"`
	Attributes entities.Attributes `json:"attributes" binding:"required"`
}
