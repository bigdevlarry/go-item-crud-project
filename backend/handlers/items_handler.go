package handlers

import (
	"errors"
	"go-test/backend/helpers"
	"go-test/backend/models/dto"
	"go-test/backend/models/entities"
	"go-test/backend/store"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemsHandler struct {
	storage store.ItemsStorage
}

func NewItemsHandler(storage store.ItemsStorage) *ItemsHandler {
	return &ItemsHandler{
		storage: storage,
	}
}

func (h *ItemsHandler) GetAll(c *gin.Context) {
	limitStr := c.Query("limit")
	query := c.Query("query")
	limit := 10 // Default limit

	hasFilters := query != ""

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			helpers.Error(c, http.StatusInternalServerError, "Invalid limit parameter")
			return
		}

		if limit < 0 {
			helpers.Error(c, http.StatusInternalServerError, "Invalid limit value")
			return
		}
	}

	items, err := h.storage.GetAll()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	var filteredItems []entities.Item
	for _, item := range items {
		// If no query, include all items
		if query == "" {
			filteredItems = append(filteredItems, item)
			continue
		}

		// use query param
		queryLower := strings.ToLower(query)
		matches := strings.Contains(strings.ToLower(item.GUID), queryLower) ||
			strings.Contains(strings.ToLower(string(item.Type)), queryLower) ||
			strings.Contains(strings.ToLower(string(item.Status)), queryLower)

		if matches {
			filteredItems = append(filteredItems, item)
		}
	}

	if hasFilters && len(filteredItems) == 0 {
		helpers.Respond(c, http.StatusOK, []any{})
		return
	}

	if limit < len(filteredItems) {
		filteredItems = filteredItems[:limit]
	}

	helpers.Respond(c, http.StatusOK, filteredItems)
}

func (h *ItemsHandler) GetByGUID(c *gin.Context) {
	guid := c.Param("guid")

	item, err := h.storage.GetByGUID(guid)
	if errors.Is(err, store.ErrNotFound) {
		helpers.Error(c, http.StatusNotFound, "Item not found")
		return
	} else if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Respond(c, http.StatusOK, *item)
}

// Create creates a new item
func (h *ItemsHandler) Create(c *gin.Context) {
	var createDTO dto.ItemCreateDTO

	if err := c.ShouldBindJSON(&createDTO); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	helpers.NormalizeCreateDTO(&createDTO)

	existingItems, err := h.storage.GetAll()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Convert DTO to Item, so we can add system fields
	item := &entities.Item{
		GUID:       uuid.New().String(),
		Amount:     createDTO.Amount,
		Type:       createDTO.Type,
		Status:     createDTO.Status,
		Attributes: createDTO.Attributes,
		Created:    time.Now(),
		Index:      len(existingItems) + 1,
	}

	if err := h.storage.Create(item); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Respond(c, http.StatusCreated, *item)
}

// Update updates an existing item
func (h *ItemsHandler) Update(c *gin.Context) {
	guid := c.Param("guid")

	existingItem, err := h.storage.GetByGUID(guid)
	if errors.Is(err, store.ErrNotFound) {
		helpers.Error(c, http.StatusNotFound, "Item not found")
		return
	} else if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	var updateDTO dto.ItemUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		helpers.ValidationError(c, err)
		return
	}

	helpers.NormalizeUpdateDTO(&updateDTO)

	updatedItem := *existingItem
	if updateDTO.Amount != nil {
		updatedItem.Amount = *updateDTO.Amount
	}
	if updateDTO.Type != nil {
		updatedItem.Type = *updateDTO.Type
	}
	if updateDTO.Status != nil {
		updatedItem.Status = *updateDTO.Status
	}
	if updateDTO.Attributes != nil {
		updatedItem.Attributes = *updateDTO.Attributes
	}

	// Update the item
	if err := h.storage.Update(&updatedItem); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Respond(c, http.StatusOK, updatedItem)
}

// Delete deletes an item
func (h *ItemsHandler) Delete(c *gin.Context) {
	guid := c.Param("guid")

	_, err := h.storage.GetByGUID(guid)
	if errors.Is(err, store.ErrNotFound) {
		helpers.Error(c, http.StatusNotFound, "Item not found")
		return
	} else if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.storage.Delete(guid)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Respond(c, http.StatusOK, []any{})
}
