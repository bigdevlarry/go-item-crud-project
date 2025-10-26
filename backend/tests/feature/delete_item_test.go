package feature

import (
	"go-test/backend/models/entities"
	"go-test/backend/models/enums"
	"go-test/backend/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCanDeleteAnItem(t *testing.T) {
	r, s := tests.SetupReadRouter()

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
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("It returns 422 when GUID is empty", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodDelete, "/items/%20", nil) // URL encoded spaces
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "GUID is required")
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

	t.Run("It returns 404 when deleting an already deleted item", func(t *testing.T) {
		// Create a new item for this test
		testItem := &entities.Item{
			GUID:   "test-guid-delete-twice",
			Amount: 100,
			Type:   enums.ADMISSION,
			Status: enums.ACCEPTED,
		}
		s.Create(testItem)

		// First delete should succeed
		req := httptest.NewRequest(http.MethodDelete, "/items/"+testItem.GUID, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)

		// Delete the same item again should return 404
		w2 := httptest.NewRecorder()
		r.ServeHTTP(w2, req)
		assert.Equal(t, http.StatusNotFound, w2.Code)
		assert.Contains(t, w2.Body.String(), "Item not found")
	})
}
