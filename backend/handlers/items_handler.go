package handlers

import (
	"errors"
	"go-test/backend/domain/dto"
	"go-test/backend/helpers"
	"go-test/backend/store"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
	query := strings.TrimSpace(c.Query("query"))

	limit, err := helpers.ParseLimit(c.Query("limit"), 10)
	if err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.storage.GetAllFiltered(query, limit)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Respond(c, http.StatusOK, items)
}

func (h *ItemsHandler) GetByGUID(c *gin.Context) {
	guid := c.Param("guid")

	if strings.TrimSpace(guid) == "" {
		helpers.Error(c, http.StatusUnprocessableEntity, "GUID is required")
		return
	}

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
		helpers.ValidationErrorResponse(c, err)
		return
	}

	count, err := h.storage.Count()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	item := helpers.NewItemFromDTO(createDTO, count+1)

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
		helpers.ValidationErrorResponse(c, err)
		return
	}

	helpers.ApplyUpdate(existingItem, updateDTO)

	// Update the item
	if err := h.storage.Update(existingItem); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.Respond(c, http.StatusOK, *existingItem)
}

// Delete deletes an item
func (h *ItemsHandler) Delete(c *gin.Context) {
	guid := c.Param("guid")

	if strings.TrimSpace(guid) == "" {
		helpers.Error(c, http.StatusUnprocessableEntity, "GUID is required")
		return
	}

	err := h.storage.Delete(guid)
	if errors.Is(err, store.ErrNotFound) {
		helpers.Error(c, http.StatusNotFound, "Item not found")
		return
	} else if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.NoContent(c, http.StatusNoContent)
}
