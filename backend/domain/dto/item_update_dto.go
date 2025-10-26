package dto

import (
	"go-test/backend/domain/enums"
	"go-test/backend/domain/models"
)

type ItemUpdateDTO struct {
	Amount     *float64           `json:"amount,omitempty" binding:"omitempty,gt=0"`
	Type       *enums.ItemType    `json:"type,omitempty" binding:"omitempty,itemtype"`
	Status     *enums.ItemStatus  `json:"status,omitempty" binding:"omitempty,itemstatus"`
	Attributes *models.Attributes `json:"attributes,omitempty" binding:"omitempty"`
}
