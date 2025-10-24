package validators

import (
	"go-test/internal/models/enums"

	"github.com/go-playground/validator/v10"
)

var validItemTypes = map[enums.ItemType]bool{
	enums.ADMISSION:  true,
	enums.SUBMISSION: true,
	enums.REVERSAL:   true,
}

var validItemStatuses = map[enums.ItemStatus]bool{
	enums.ACCEPTED: true,
	enums.DECLINED: true,
}

func ValidateItemType(fl validator.FieldLevel) bool {
	itemType := enums.ItemType(fl.Field().String())
	_, ok := validItemTypes[itemType]
	return ok
}

func ValidateItemStatus(fl validator.FieldLevel) bool {
	itemStatus := enums.ItemStatus(fl.Field().String())
	_, ok := validItemStatuses[itemStatus]
	return ok
}
