package tests

import (
	"go-test/backend/models/entities"
	"go-test/backend/models/enums"
	"go-test/backend/tests/helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCanCreateAnItem(t *testing.T) {
	r, _ := helper.SetupReadRouter()

	t.Run("It can create an item with valid data", func(t *testing.T) {
		// Arrange
		validPayload := createValidCreatePayload()
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(validPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), `"amount":100`)
		assert.Contains(t, w.Body.String(), `"type":"ADMISSION"`)
		assert.Contains(t, w.Body.String(), `"status":"ACCEPTED"`)
		assert.Contains(t, w.Body.String(), `"first_name":"John"`)
		assert.Contains(t, w.Body.String(), `"last_name":"Doe"`)
		assert.Contains(t, w.Body.String(), `"first_name":"Jane"`)
		assert.Contains(t, w.Body.String(), `"last_name":"Smith"`)
		assert.Contains(t, w.Body.String(), `"sort_code":"123456"`)
		assert.Contains(t, w.Body.String(), `"account_number":"12345678"`)
		assert.Contains(t, w.Body.String(), `"sort_code":"876543"`)
		assert.Contains(t, w.Body.String(), `"account_number":"87654321"`)
		assert.Contains(t, w.Body.String(), `"guid"`)
		assert.Contains(t, w.Body.String(), `"created"`)
	})

	t.Run("It can create multiple items", func(t *testing.T) {
		// Arrange
		validPayload := createValidCreatePayload()
		req1 := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(validPayload))
		req1.Header.Set("Content-Type", "application/json")
		w1 := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w1, req1)
		assert.Equal(t, http.StatusCreated, w1.Code)

		req2 := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(validPayload))
		req2.Header.Set("Content-Type", "application/json")
		w2 := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w2, req2)

		// Assert
		assert.Equal(t, http.StatusCreated, w2.Code)
	})

	t.Run("It returns 400 error when required fields are missing", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Amount")
		assert.Contains(t, w.Body.String(), "Type")
		assert.Contains(t, w.Body.String(), "Status")
		assert.Contains(t, w.Body.String(), "Attributes")
	})

	t.Run("It returns 400 error when amount is not greater than 0", func(t *testing.T) {
		// Arrange
		invalidPayload := `{
			"amount": 0,
			"type": "ADMISSION",
			"status": "ACCEPTED",
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
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Amount")
		assert.Contains(t, w.Body.String(), "required")
	})

	t.Run("It returns 400 error when amount is negative", func(t *testing.T) {
		// Arrange
		invalidPayload := `{
			"amount": -100,
			"type": "ADMISSION",
			"status": "ACCEPTED",
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
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Amount")
		assert.Contains(t, w.Body.String(), "gt")
	})

	t.Run("It returns 400 error when type is invalid", func(t *testing.T) {
		// Arrange
		invalidPayload := `{
			"amount": 100,
			"type": "INVALID_TYPE",
			"status": "ACCEPTED",
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
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Type")
	})

	t.Run("It returns 400 error when status is invalid", func(t *testing.T) {
		// Arrange
		invalidPayload := `{
			"amount": 100,
			"type": "ADMISSION",
			"status": "INVALID_STATUS",
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
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Status")
	})

	t.Run("It returns 400 error when attributes are missing", func(t *testing.T) {
		// Arrange
		invalidPayload := `{
			"amount": 100,
			"type": "ADMISSION",
			"status": "ACCEPTED"
		}`
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Attributes")
	})

	t.Run("It returns 400 error when debtor information is missing", func(t *testing.T) {
		invalidPayload := `{
			"amount": 100,
			"type": "ADMISSION",
			"status": "ACCEPTED",
			"attributes": {
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

		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Debtor")
	})

	t.Run("It returns 400 error when beneficiary information is missing", func(t *testing.T) {
		invalidPayload := `{
			"amount": 100,
			"type": "ADMISSION",
			"status": "ACCEPTED",
			"attributes": {
				"debtor": {
					"first_name": "John",
					"last_name": "Doe",
					"account": {
						"sort_code": "123456",
						"account_number": "12345678"
					}
				}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Beneficiary")
	})

	t.Run("It returns 400 error when account information is missing", func(t *testing.T) {
		invalidPayload := `{
			"amount": 100,
			"type": "ADMISSION",
			"status": "ACCEPTED",
			"attributes": {
				"debtor": {
					"first_name": "John",
					"last_name": "Doe"
				},
				"beneficiary": {
					"first_name": "Jane",
					"last_name": "Smith"
				}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Account")
	})

	t.Run("It returns 400 error when JSON is malformed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{invalid json`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

	t.Run("It returns 400 error when request body is empty", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/items", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
}

func TestItAcceptsCaseInsensitiveEnums(t *testing.T) {
	r, _ := helper.SetupReadRouter()

	t.Run("It accepts lowercase enum values", func(t *testing.T) {
		// Arrange
		payload := `{
			"amount": 100,
			"type": "admission",
			"status": "accepted",
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
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), `"type":"ADMISSION"`)
		assert.Contains(t, w.Body.String(), `"status":"ACCEPTED"`)
	})

	t.Run("It accepts mixed case enum values", func(t *testing.T) {
		// Arrange
		payload := `{
			"amount": 100,
			"type": "SuBmIsSiOn",
			"status": "DeClInEd",
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
		req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Act
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), `"type":"SUBMISSION"`)
		assert.Contains(t, w.Body.String(), `"status":"DECLINED"`)
	})
}

func createValidCreatePayload() string {
	return `{
		"amount": 100,
		"type": "ADMISSION",
		"status": "ACCEPTED",
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

func TestStorageValidationForCreate(t *testing.T) {
	_, s := helper.SetupReadRouter()

	t.Run("It returns error when creating item with nil pointer", func(t *testing.T) {
		// Act
		err := s.Create(nil)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "item cannot be nil")
	})

	t.Run("It returns error when creating item with empty GUID", func(t *testing.T) {
		// Arrange
		item := &entities.Item{
			GUID:   "", // Empty GUID
			Amount: 100,
			Type:   enums.ADMISSION,
			Status: enums.ACCEPTED,
		}

		// Act
		err := s.Create(item)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "item GUID cannot be empty")
	})
}
