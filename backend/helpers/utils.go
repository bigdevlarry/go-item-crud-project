package helpers

import (
	"errors"
	"go-test/backend/models/dto"
	"go-test/backend/models/entities"
	"go-test/backend/models/enums"
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

func NewItemFromDTO(dto dto.ItemCreateDTO, index int) *entities.Item {
	return &entities.Item{
		GUID:       uuid.New().String(),
		Amount:     dto.Amount,
		Type:       enums.ItemType(strings.ToUpper(string(dto.Type))),
		Status:     enums.ItemStatus(strings.ToUpper(string(dto.Status))),
		Attributes: dto.Attributes,
		Created:    time.Now(),
		Index:      index,
	}
}

func ApplyUpdate(item *entities.Item, dto dto.ItemUpdateDTO) {
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

func ApplyLimit(items []entities.Item, limit int) []entities.Item {
	if limit <= 0 || limit >= len(items) {
		return items
	}
	return items[:limit]
}

func CopyItems(m map[string]entities.Item) []entities.Item {
	items := make([]entities.Item, 0, len(m))
	for _, item := range m {
		items = append(items, item)
	}
	return items
}
