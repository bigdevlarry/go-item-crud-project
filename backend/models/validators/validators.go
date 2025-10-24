package validators

import (
	"go-test/backend/models/enums"
	"regexp"
	"strings"

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
	// handle case-insensitive validation
	itemType := enums.ItemType(strings.ToUpper(fl.Field().String()))
	_, ok := validItemTypes[itemType]
	return ok
}

func ValidateItemStatus(fl validator.FieldLevel) bool {
	// handle case-insensitive validation
	itemStatus := enums.ItemStatus(strings.ToUpper(fl.Field().String()))
	_, ok := validItemStatuses[itemStatus]
	return ok
}

// ValidateSortCode validates sort code format (00-00-00)
func ValidateSortCode(fl validator.FieldLevel) bool {
	sortCode := fl.Field().String()
	sortCodeRegex := regexp.MustCompile(`^\d{2}-\d{2}-\d{2}$`)
	return sortCodeRegex.MatchString(sortCode)
}
