package dto

import (
	"go-test/internal/models/entities"
	"go-test/internal/models/enums"
)

type ItemUpdateDTO struct {
	Amount     *float64             `json:"amount,omitempty" binding:"omitempty,gt=0"`
	Type       *enums.ItemType      `json:"type,omitempty" binding:"omitempty,itemtype"`
	Status     *enums.ItemStatus    `json:"status,omitempty" binding:"omitempty,itemstatus"`
	Attributes *entities.Attributes `json:"attributes,omitempty" binding:"omitempty"`
}
