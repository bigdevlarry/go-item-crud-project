package store

import (
	"errors"
	"go-test/backend/models/entities"
	"strings"
)

var ErrNotFound = errors.New("item not found")

type ItemsStorage interface {
	GetAll() ([]entities.Item, error)
	GetAllFiltered(query string, limit int) ([]entities.Item, error)
	GetByGUID(guid string) (*entities.Item, error)
	Count() (int, error)
	Create(item *entities.Item) error
	Update(item *entities.Item) error
	Delete(guid string) error
}

type ItemsStore struct {
	items map[string]entities.Item
}

// NewStore creates a new, thread-safe in-memory item store
func NewStore() *ItemsStore {
	return &ItemsStore{
		items: make(map[string]entities.Item),
	}
}

// GetAll returns all items
func (is *ItemsStore) GetAll() ([]entities.Item, error) {
	items := make([]entities.Item, 0, len(is.items))
	for _, item := range is.items {
		items = append(items, item)
	}
	return items, nil
}

// GetAllFiltered returns filtered and limited items
func (is *ItemsStore) GetAllFiltered(query string, limit int) ([]entities.Item, error) {
	filteredItems := make([]entities.Item, 0)

	// If no query, return all items (with limit applied)
	if query == "" {
		items := make([]entities.Item, 0, len(is.items))
		for _, item := range is.items {
			items = append(items, item)
		}

		// Apply limit
		if limit > 0 {
			if limit > len(items) {
				limit = len(items)
			}
			items = items[:limit]
		}
		return items, nil
	}

	// Filter items based on query
	queryLower := strings.ToLower(query)
	for _, item := range is.items {
		matches := strings.Contains(strings.ToLower(item.GUID), queryLower) ||
			strings.Contains(strings.ToLower(string(item.Type)), queryLower) ||
			strings.Contains(strings.ToLower(string(item.Status)), queryLower)

		if matches {
			filteredItems = append(filteredItems, item)
		}
	}

	// Apply limit
	if limit > 0 {
		if limit > len(filteredItems) {
			limit = len(filteredItems)
		}
		filteredItems = filteredItems[:limit]
	}

	return filteredItems, nil
}

// GetByGUID returns an item by GUID
func (is *ItemsStore) GetByGUID(guid string) (*entities.Item, error) {
	item, exists := is.items[guid]
	if !exists {
		return nil, ErrNotFound
	}
	return &item, nil
}

// Count returns the total number of items
func (is *ItemsStore) Count() (int, error) {
	return len(is.items), nil
}

// Create adds a new item
func (is *ItemsStore) Create(item *entities.Item) error {
	is.items[item.GUID] = *item
	return nil
}

// Update updates an item by a given GUID
func (is *ItemsStore) Update(item *entities.Item) error {
	is.items[item.GUID] = *item
	return nil
}

// Delete removes an item by GUID
func (is *ItemsStore) Delete(guid string) error {
	_, exists := is.items[guid]
	if !exists {
		return ErrNotFound
	}
	delete(is.items, guid)
	return nil
}
