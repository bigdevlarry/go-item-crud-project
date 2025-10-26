package helpers

import (
	"errors"
	"go-test/backend/domain/dto"
	"go-test/backend/domain/enums"
	"go-test/backend/domain/models"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func ParseLimit(limitStr string, defaultLimit int) (int, error) {
	if limitStr == "" {
		return defaultLimit, nil
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, errors.New("invalid limit parameter")
	}

	if limit < 0 {
		return 0, errors.New("invalid limit value")
	}

	return limit, nil
}

func NewItemFromDTO(dto dto.ItemCreateDTO, index int) *models.Item {
	return &models.Item{
		GUID:       uuid.New().String(),
		Amount:     dto.Amount,
		Type:       enums.ItemType(strings.ToUpper(string(dto.Type))),
		Status:     enums.ItemStatus(strings.ToUpper(string(dto.Status))),
		Attributes: dto.Attributes,
		Created:    time.Now(),
		Index:      index,
	}
}

func ApplyUpdate(item *models.Item, dto dto.ItemUpdateDTO) {
	if dto.Amount != nil {
		item.Amount = *dto.Amount
	}
	if dto.Type != nil {
		item.Type = enums.ItemType(strings.ToUpper(string(*dto.Type)))
	}
	if dto.Status != nil {
		item.Status = enums.ItemStatus(strings.ToUpper(string(*dto.Status)))
	}
	if dto.Attributes != nil {
		item.Attributes = *dto.Attributes
	}
}

func ApplyLimit(items []models.Item, limit int) []models.Item {
	if limit <= 0 || limit >= len(items) {
		return items
	}
	return items[:limit]
}

func CopyItems(m map[string]models.Item) []models.Item {
	items := make([]models.Item, 0, len(m))
	for _, item := range m {
		items = append(items, item)
	}

	// Sort by Index for consistency on the FE
	sort.Slice(items, func(i, j int) bool {
		return items[i].Index < items[j].Index
	})

	return items
}
