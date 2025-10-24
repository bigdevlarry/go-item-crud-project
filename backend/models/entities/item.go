package entities

import (
	"go-test/backend/models/enums"
	"time"
)

type Item struct {
	GUID       string           `json:"guid"`
	Index      int              `json:"index"`
	Amount     float64          `json:"amount" binding:"required,gt=0"`
	Type       enums.ItemType   `json:"type" binding:"required,itemtype"`
	Status     enums.ItemStatus `json:"status" binding:"required,itemstatus"`
	Created    time.Time        `json:"created"`
	Attributes Attributes       `json:"attributes" binding:"required"`
}
type Attributes struct {
	Debtor      Party `json:"debtor" binding:"required"`
	Beneficiary Party `json:"beneficiary" binding:"required"`
}
