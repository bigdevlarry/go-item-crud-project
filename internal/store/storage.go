package store

import (
	"go-test/internal/models"
)

type ItemsStorage interface {
	GetAll() ([]models.Item, error)
	GetByGUID(guid string) (*models.Item, error)
	Create(item *models.Item) error
	Update(item *models.Item) error
	Delete(guid string) error
}

type ItemsStore struct {
	items map[string]models.Item
}

// NewStore creates a new items store
func NewStore() *ItemsStore {
	return &ItemsStore{items: make(map[string]models.Item)}
}

// GetAll returns all items
func (is *ItemsStore) GetAll() ([]models.Item, error) {
	items := make([]models.Item, 0, len(is.items))
	for _, item := range is.items {
		items = append(items, item)
	}
	return items, nil
}

// GetByGUID returns an item by GUID
func (is *ItemsStore) GetByGUID(guid string) (*models.Item, error) {
	item, exists := is.items[guid]
	if !exists {
		return nil, nil // Item not found
	}
	return &item, nil
}

// Create adds a new item
func (is *ItemsStore) Create(item *models.Item) error {
	is.items[item.GUID] = *item
	return nil
}

// Update updates an item by a given GUID
func (is *ItemsStore) Update(item *models.Item) error {
	is.items[item.GUID] = *item
	return nil
}

// Delete removes an item by GUID
func (is *ItemsStore) Delete(guid string) error {
	delete(is.items, guid)
	return nil
}
