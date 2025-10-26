import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useItemsStore } from '../items'
import type { Item, ItemCreateDTO, ItemUpdateDTO } from '@/types'

// Mock fetch globally
global.fetch = vi.fn()

describe('Items Store', () => {
  let store: ReturnType<typeof useItemsStore>

  beforeEach(() => {
    // Create a fresh Pinia instance
    const pinia = createPinia()
    setActivePinia(pinia)

    // Get the store instance
    store = useItemsStore()

    // Clear all mocks
    vi.clearAllMocks()
  })

  describe('fetchItems', () => {
    it('fetches items successfully', async () => {
      const mockItems: Item[] = [
        {
          guid: '1',
          index: 1,
          amount: 100.50,
          type: 'ADMISSION',
          status: 'ACCEPTED',
          created: '2024-01-15T10:30:00Z',
          attributes: {
            debtor: {
              first_name: 'John',
              last_name: 'Doe',
              account: {
                sort_code: '12-34-56',
                account_number: '12345678'
              }
            },
            beneficiary: {
              first_name: 'Jane',
              last_name: 'Smith',
              account: {
                sort_code: '78-90-12',
                account_number: '87654321'
              }
            }
          }
        }
      ]

      // Mock successful response
      ;(global.fetch as any).mockResolvedValueOnce({
        ok: true,
        json: async () => mockItems
      })

      await store.fetchItems()

      expect(store.items).toEqual(mockItems)
      expect(store.loading).toBe(false)
      expect(store.error).toBe(null)
    })

    it('fetches items with search query', async () => {
      const mockItems: Item[] = []

      ;(global.fetch as any).mockResolvedValueOnce({
        ok: true,
        json: async () => mockItems
      })

      await store.fetchItems('test query')

      expect(global.fetch).toHaveBeenCalledWith(
        'http://localhost:8080/items?query=test%20query',
        undefined
      )
    })
  })

  describe('createItem', () => {
    it('creates item successfully', async () => {
      const newItemData: ItemCreateDTO = {
        amount: 200.75,
        type: 'SUBMISSION',
        status: 'ACCEPTED',
        attributes: {
          debtor: {
            first_name: 'John',
            last_name: 'Doe',
            account: {
              sort_code: '12-34-56',
              account_number: '12345678'
            }
          },
          beneficiary: {
            first_name: 'Jane',
            last_name: 'Smith',
            account: {
              sort_code: '78-90-12',
              account_number: '87654321'
            }
          }
        }
      }

      const createdItem: Item = {
        guid: 'new-guid',
        index: 1,
        ...newItemData,
        created: '2024-01-15T10:30:00Z'
      }

      // Mock successful response
      ;(global.fetch as any).mockResolvedValueOnce({
        ok: true,
        json: async () => createdItem
      })

      const result = await store.createItem(newItemData)

      expect(result).toEqual(createdItem)
      expect(store.items).toHaveLength(1)
      expect(store.items[0]).toEqual(createdItem)
      expect(store.error).toBe(null)
    })
  })

  describe('updateItem', () => {
    it('updates item successfully', async () => {
      const existingItem: Item = {
        guid: 'test-guid',
        index: 1,
        amount: 100.50,
        type: 'ADMISSION',
        status: 'ACCEPTED',
        created: '2024-01-15T10:30:00Z',
        attributes: {
          debtor: {
            first_name: 'John',
            last_name: 'Doe',
            account: {
              sort_code: '12-34-56',
              account_number: '12345678'
            }
          },
          beneficiary: {
            first_name: 'Jane',
            last_name: 'Smith',
            account: {
              sort_code: '78-90-12',
              account_number: '87654321'
            }
          }
        }
      }

      // Add item to store first
      store.items.push(existingItem)

      const updateData: ItemUpdateDTO = {
        amount: 150.75
      }

      const updatedItem: Item = {
        ...existingItem,
        amount: 150.75
      }

      // Mock successful response
      ;(global.fetch as any).mockResolvedValueOnce({
        ok: true,
        json: async () => updatedItem
      })

      const result = await store.updateItem('test-guid', updateData)

      expect(result).toEqual(updatedItem)
      expect(store.items[0]?.amount).toBe(150.75)
      expect(store.error).toBe(null)
    })
  })

  describe('deleteItem', () => {
    it('deletes item successfully', async () => {
      const itemToDelete: Item = {
        guid: 'test-guid',
        index: 1,
        amount: 100.50,
        type: 'ADMISSION',
        status: 'ACCEPTED',
        created: '2024-01-15T10:30:00Z',
        attributes: {
          debtor: {
            first_name: 'John',
            last_name: 'Doe',
            account: {
              sort_code: '12-34-56',
              account_number: '12345678'
            }
          },
          beneficiary: {
            first_name: 'Jane',
            last_name: 'Smith',
            account: {
              sort_code: '78-90-12',
              account_number: '87654321'
            }
          }
        }
      }

      // Add item to store first
      store.items.push(itemToDelete)

      // Mock successful response
      ;(global.fetch as any).mockResolvedValueOnce({
        ok: true,
        json: async () => ({})
      })

      const result = await store.deleteItem('test-guid')

      expect(result).toBe(true)
      expect(store.items).not.toContain(itemToDelete)
      expect(store.error).toBe(null)
    })
  })

  describe('setSearchQuery', () => {
    it('sets search query', () => {
      store.setSearchQuery('test query')
      expect(store.searchQuery).toBe('test query')
    })
  })
})
