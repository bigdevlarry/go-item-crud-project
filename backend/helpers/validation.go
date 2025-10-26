package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Errors map[string]string `json:"errors"`
}

func ValidationErrorResponse(c *gin.Context, err error) {
	validationErrors := ValidationError{
		Errors: make(map[string]string),
	}

	// Handle JSON parsing (type) errors
	var jsonTypeErr *json.UnmarshalTypeError
	if ok := errors.As(err, &jsonTypeErr); ok {
		fieldPath := strings.ToLower(jsonTypeErr.Field)
		expected := jsonTypeErr.Type.String()
		switch expected {
		case "float64":
			validationErrors.Errors[fieldPath] = "This field must be a number"
		default:
			validationErrors.Errors[fieldPath] = "Invalid data type"
		}
	}

	// Handle validator-based struct validation errors
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		for _, fieldErr := range validationErr {
			fieldPath := strings.ToLower(fieldErr.Field())

			var message string
			switch fieldErr.Tag() {
			case "required":
				message = "This field is required"
			case "gt":
				message = "Value must be greater than 0"
			case "len":
				message = "Must be exactly 8 digits"
			case "sortcode":
				message = "Sort code must be in the format 00-00-00"
			case "itemtype":
				message = "Invalid item type. Must be ADMISSION, SUBMISSION, or REVERSAL"
			case "itemstatus":
				message = "Invalid item status. Must be ACCEPTED or DECLINED"
			default:
				message = "Invalid value"
			}

			validationErrors.Errors[fieldPath] = message
		}
	}

	c.JSON(http.StatusBadRequest, validationErrors)
}
