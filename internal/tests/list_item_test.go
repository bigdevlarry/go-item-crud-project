package tests

import (
	"go-test/internal/models/entities"
	"go-test/internal/models/enums"
	"go-test/internal/tests/helper"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemsListAndFilter(t *testing.T) {
	r, s := helper.SetupReadRouter()

	item := &entities.Item{
		GUID:   "test-guid-123",
		Amount: 100,
		Type:   enums.ADMISSION,
		Status: enums.ACCEPTED,
	}
	s.Create(item)

	t.Run("It can list all items", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can apply limit parameter to items", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can filter items by GUID", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?guid=test-guid&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can filter items by type", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?type=ADMISSION&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can filter items by status", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?status=ACCEPTED&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can filter items by type and status", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?type=ADMISSION&status=ACCEPTED&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It returns empty results when no items match GUID filter", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?guid=nonexistent-guid", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "No items found")
	})

	t.Run("It returns empty results when no items match status filter", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?status=NONEXISTENT_STATUS", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "No items found")
	})

	t.Run("It returns empty results when no items match type filter", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?type=NONEXISTENT_TYPE", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "No items found")
	})

	t.Run("It returns empty results when no items match multiple filters", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?type=ADMISSION&status=DECLINED", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "No items found")
	})

	t.Run("It returns a server error with invalid query params", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?limit=invalid", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid limit parameter")
	})

	t.Run("It returns a server error with negative query params value", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?limit=-1000", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid limit value")
	})
}
