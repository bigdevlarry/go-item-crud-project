package tests

import (
	"go-test/internal/models"
	"go-test/internal/models/enums"
	"go-test/internal/tests/helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCanUpdateAnItem(t *testing.T) {
	r, s := helper.SetupReadRouter()

	item := &models.Item{
		GUID:   "test-guid-123",
		Amount: 100,
		Type:   enums.ADMISSION,
		Status: enums.ACCEPTED,
	}
	s.Create(item)

	t.Run("It can successfully update an item", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodPut, "/items/"+item.GUID, strings.NewReader(createUpdatePayload()))
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"guid":"test-guid-123"`)
		assert.Contains(t, w.Body.String(), `"amount":200`)
		assert.Contains(t, w.Body.String(), `"type":"SUBMISSION"`)
		assert.Contains(t, w.Body.String(), `"status":"DECLINED"`)
		assert.Contains(t, w.Body.String(), `"first_name":"John"`)
		assert.Contains(t, w.Body.String(), `"last_name":"Doe"`)
		assert.Contains(t, w.Body.String(), `"first_name":"Jane"`)
		assert.Contains(t, w.Body.String(), `"last_name":"Smith"`)
		assert.Contains(t, w.Body.String(), `"sort_code":"123456"`)
		assert.Contains(t, w.Body.String(), `"account_number":"12345678"`)
		assert.Contains(t, w.Body.String(), `"sort_code":"876543"`)
		assert.Contains(t, w.Body.String(), `"account_number":"87654321"`)
	})

	t.Run("It returns 404 error when an item does not exist", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodPut, "/items/nonexistent-guid", strings.NewReader(createUpdatePayload()))
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Item not found")
	})

	t.Run("It returns 400 error when amount validation fails", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodPut, "/items/"+item.GUID, strings.NewReader(`{"amount": 0}`))
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Field validation for 'Amount' failed on the 'gt'")
	})

	t.Run("It returns 400 error when the status or type is invalid", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodPut, "/items/"+item.GUID, strings.NewReader(`{"type": "invalid", "status": "pending"}`))
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Field validation for 'Status' failed on the 'itemstatus'")
		assert.Contains(t, w.Body.String(), "Field validation for 'Type' failed on the 'itemtype'")
	})
}

func createUpdatePayload() string {
	return `{
        "amount": 200,
        "type": "SUBMISSION",
        "status": "DECLINED",
        "attributes": {
            "debtor": {
                "first_name": "John",
                "last_name": "Doe",
                "account": {
                    "sort_code": "123456",
                    "account_number": "12345678"
                }
            },
            "beneficiary": {
                "first_name": "Jane",
                "last_name": "Smith",
                "account": {
                    "sort_code": "876543",
                    "account_number": "87654321"
                }
            }
        }
    }`
}
