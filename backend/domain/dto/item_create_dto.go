package dto

import (
	"go-test/backend/domain/enums"
	"go-test/backend/domain/models"
)

type ItemCreateDTO struct {
	Amount     float64           `json:"amount" binding:"required,gt=0"`
	Type       enums.ItemType    `json:"type" binding:"required,itemtype"`
	Status     enums.ItemStatus  `json:"status" binding:"required,itemstatus"`
	Attributes models.Attributes `json:"attributes" binding:"required"`
}
