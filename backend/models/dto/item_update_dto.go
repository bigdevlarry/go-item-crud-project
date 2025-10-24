package dto

import (
	"go-test/backend/models/entities"
	"go-test/backend/models/enums"
)

type ItemUpdateDTO struct {
	Amount     *float64             `json:"amount,omitempty" binding:"omitempty,gt=0"`
	Type       *enums.ItemType      `json:"type,omitempty" binding:"omitempty,itemtype"`
	Status     *enums.ItemStatus    `json:"status,omitempty" binding:"omitempty,itemstatus"`
	Attributes *entities.Attributes `json:"attributes,omitempty" binding:"omitempty"`
}
