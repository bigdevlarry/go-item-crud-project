package tests

import (
	"go-test/backend/models/entities"
	"go-test/backend/models/enums"
	"go-test/backend/tests/helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCanDeleteAnItem(t *testing.T) {
	r, s := helper.SetupReadRouter()

	item := &entities.Item{
		GUID:   "test-guid-123",
		Amount: 100,
		Type:   enums.ADMISSION,
		Status: enums.ACCEPTED,
	}
	s.Create(item)

	t.Run("It can delete an item", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodDelete, "/items/"+item.GUID, nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Item deleted successfully")
	})

	t.Run("It returns 404 when an item does not exist", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodDelete, "/items/nonexistent-guid", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Item not found")
	})
}
