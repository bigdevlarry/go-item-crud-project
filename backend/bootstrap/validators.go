package bootstrap

import (
	"go-test/backend/domain/validators"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// RegisterCustomValidators registers all custom validation functions
func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("itemtype", validators.ValidateItemType)
		if err != nil {
			log.Fatal("Failed to register itemtype validator:", err)
		}

		err = v.RegisterValidation("itemstatus", validators.ValidateItemStatus)
		if err != nil {
			log.Fatal("Failed to register itemstatus validator:", err)
		}

		err = v.RegisterValidation("sortcode", validators.ValidateSortCode)
		if err != nil {
			log.Fatal("Failed to register sortcode validator:", err)
		}
	} else {
		log.Fatal("Failed to get validator engine")
	}
}
