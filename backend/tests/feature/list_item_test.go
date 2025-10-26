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

func TestItemsListAndFilter(t *testing.T) {
	r, s := tests.SetupReadRouter()

	item := &models.Item{
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

	t.Run("It can filter items by GUID using global search", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=test-guid&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can filter items by type using global search", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=ADMISSION&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can filter items by status using global search", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=ACCEPTED&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It can filter items by partial match using global search", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=test&limit=5", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("It returns empty results when no items match query", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=nonexistent", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "[]", w.Body.String())
	})

	t.Run("It returns empty results when query matches no items", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=NONEXISTENT", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "[]", w.Body.String())
	})

	t.Run("It returns a server error with invalid query params", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?limit=invalid", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid limit parameter")
	})

	t.Run("It returns a server error with negative query params value", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?limit=-1000", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid limit value")
	})

	t.Run("It returns all items when limit is zero or not provided", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?limit=0", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), item.GUID)
	})

	t.Run("It can filter items regardless of case sensitivity", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=adMiSsIon", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "ADMISSION")
	})

	t.Run("It trims whitespace in query before filtering", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/items?query=%20ADMISSION%20", nil)
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "ADMISSION")
	})
}
