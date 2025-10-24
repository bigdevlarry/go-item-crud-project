package handlers

import (
	"go-test/internal/models/dto"
	"go-test/internal/models/entities"
	"go-test/internal/models/enums"
	"go-test/internal/store"
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
	limit := 10 // Default limit

	searchGUID := c.Query("guid")
	searchType := c.Query("type")
	searchStatus := c.Query("status")

	hasFilters := searchGUID != "" || searchType != "" || searchStatus != ""

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid limit parameter"})
			return
		}

		if limit < 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid limit value"})
			return
		}
	}

	items, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var filteredItems []entities.Item
	for _, item := range items {
		if searchGUID != "" && !strings.Contains(item.GUID, searchGUID) {
			continue
		}
		if searchType != "" && item.Type != enums.ItemType(searchType) {
			continue
		}
		if searchStatus != "" && item.Status != enums.ItemStatus(searchStatus) {
			continue
		}

		filteredItems = append(filteredItems, item)
	}

	if hasFilters && len(filteredItems) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No items found",
			"items":   []entities.Item{},
		})
		return
	}

	if limit < len(filteredItems) {
		items = filteredItems[:limit]
	}

	c.JSON(http.StatusOK, items)
}

func (h *ItemsHandler) GetByGUID(c *gin.Context) {
	guid := c.Param("guid")

	item, err := h.storage.GetByGUID(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, []entities.Item{*item})
}

// Create creates a new item
func (h *ItemsHandler) Create(c *gin.Context) {
	var createDTO dto.ItemCreateDTO

	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingItems, err := h.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, []entities.Item{*item})
}

// Update updates an existing item
func (h *ItemsHandler) Update(c *gin.Context) {
	guid := c.Param("guid")

	// Check if item exists first
	existingItem, err := h.storage.GetByGUID(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if existingItem == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var updateDTO dto.ItemUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateDTO.Amount != nil {
		existingItem.Amount = *updateDTO.Amount
	}
	if updateDTO.Type != nil {
		existingItem.Type = *updateDTO.Type
	}
	if updateDTO.Status != nil {
		existingItem.Status = *updateDTO.Status
	}
	if updateDTO.Attributes != nil {
		existingItem.Attributes = *updateDTO.Attributes
	}

	// Update the item
	if err := h.storage.Update(existingItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, []entities.Item{*existingItem})
}

// Delete deletes an item
func (h *ItemsHandler) Delete(c *gin.Context) {
	guid := c.Param("guid")

	item, err := h.storage.GetByGUID(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	err = h.storage.Delete(guid)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item deleted successfully",
	})
}
