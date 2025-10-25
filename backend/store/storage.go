package store

import (
	"errors"
	"go-test/backend/helpers"
	"go-test/backend/models/entities"
	"strings"
	"sync"
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
	mutex sync.RWMutex
}

// NewStore creates a new, thread-safe in-memory item store
func NewStore() *ItemsStore {
	return &ItemsStore{
		items: make(map[string]entities.Item),
	}
}

// GetAll returns all items
func (is *ItemsStore) GetAll() ([]entities.Item, error) {
	is.mutex.RLock()
	defer is.mutex.RUnlock()

	items := make([]entities.Item, 0, len(is.items))
	for _, item := range is.items {
		items = append(items, item)
	}
	return items, nil
}

// GetAllFiltered returns filtered and limited items
func (is *ItemsStore) GetAllFiltered(query string, limit int) ([]entities.Item, error) {
	is.mutex.RLock()
	defer is.mutex.RUnlock()

	items := helpers.CopyItems(is.items)

	if query != "" {
		filtered := make([]entities.Item, 0)
		queryLower := strings.ToLower(query)
		for _, item := range items {
			// Pre-compute lowercase values to avoid repeated allocations
			guidLower := strings.ToLower(item.GUID)
			typeLower := strings.ToLower(string(item.Type))
			statusLower := strings.ToLower(string(item.Status))

			matches := strings.Contains(guidLower, queryLower) ||
				strings.Contains(typeLower, queryLower) ||
				strings.Contains(statusLower, queryLower)

			if matches {
				filtered = append(filtered, item)
			}
		}
		items = filtered
	}

	return helpers.ApplyLimit(items, limit), nil
}

// GetByGUID returns an item by GUID
func (is *ItemsStore) GetByGUID(guid string) (*entities.Item, error) {
	is.mutex.RLock()
	defer is.mutex.RUnlock()

	item, exists := is.items[guid]
	if !exists {
		return nil, ErrNotFound
	}
	return &item, nil
}

// Count returns the total number of items
func (is *ItemsStore) Count() (int, error) {
	is.mutex.RLock()
	defer is.mutex.RUnlock()

	return len(is.items), nil
}

// Create adds a new item
func (is *ItemsStore) Create(item *entities.Item) error {
	if item == nil {
		return errors.New("item cannot be nil")
	}
	if item.GUID == "" {
		return errors.New("item GUID cannot be empty")
	}

	is.mutex.Lock()
	defer is.mutex.Unlock()

	is.items[item.GUID] = *item
	return nil
}

// Update updates an item by a given GUID
func (is *ItemsStore) Update(item *entities.Item) error {
	if item == nil {
		return errors.New("item cannot be nil")
	}
	if item.GUID == "" {
		return errors.New("item GUID cannot be empty")
	}

	is.mutex.Lock()
	defer is.mutex.Unlock()

	_, exists := is.items[item.GUID]
	if !exists {
		return ErrNotFound
	}

	is.items[item.GUID] = *item
	return nil
}

// Delete removes an item by GUID
func (is *ItemsStore) Delete(guid string) error {
	is.mutex.Lock()
	defer is.mutex.Unlock()

	_, exists := is.items[guid]
	if !exists {
		return ErrNotFound
	}
	delete(is.items, guid)
	return nil
}
