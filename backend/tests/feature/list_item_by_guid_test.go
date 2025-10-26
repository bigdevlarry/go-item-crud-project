package feature

import (
	"go-test/backend/domain/enums"
	"go-test/backend/domain/models"
	"go-test/backend/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCanListItemByGuid(t *testing.T) {
	r, s := tests.SetupReadRouter()

	item := &models.Item{
		GUID:   "test-guid-123",
		Amount: 100,
		Type:   enums.ADMISSION,
		Status: enums.ACCEPTED,
	}
	s.Create(item)

	t.Run("It can list an item by GUID", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items/"+item.GUID, nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It returns an error when an item does not exist", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items/not-found", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Item not found")
	})

	t.Run("It returns 422 when GUID is empty", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items/%20", nil) // URL encoded space
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "GUID is required")
	})
}
