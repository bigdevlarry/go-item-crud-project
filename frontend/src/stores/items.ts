import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import type {Item, ItemCreateDTO, ItemUpdateDTO} from '@/types'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

// Request helper
async function request<T>(endpoint: string, options?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`, options)
  if (!response.ok) {
    let msg = `HTTP error! status: ${response.status}`
    try {
      const err = await response.json()
      if (err.error) msg = err.error
    } catch (_) {
      /* fallback to default error message */
    }
    throw new Error(msg)
  }
  return await response.json() as Promise<T>
}

export const useItemsStore = defineStore(
  'items', () => {
  // State
  const items = ref<Item[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const searchQuery = ref('')

  // Actions
  const setLoading = (isLoading: boolean) => {
    loading.value = isLoading
  }

  const setError = (err: string | null) => {
    error.value = err
  }

  const setSearchQuery = (query: string) => {
    searchQuery.value = query
  }

  // Fetch all items
  const fetchItems = async (query?: string) => {
    try {
      setLoading(true)
      setError(null)

      // Clear items first to prevent stale data
      items.value = []

      const endpoint = query?.trim() ? `/items?query=${encodeURIComponent(query.trim())}` : '/items'
      const data = await request<Item[]>(endpoint)

      // Handle different response formats
      if (Array.isArray(data)) {
        items.value = data
      } else {
        items.value = []
      }
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch items'
      setError(errorMessage)
      items.value = []
    } finally {
      setLoading(false)
    }
  }

  // Create new item
  const createItem = async (itemData: ItemCreateDTO): Promise<Item | null> => {
    try {
      setError(null)

      const newItem = await request<Item>('/items', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(itemData),
      })

      items.value = [...items.value, newItem]
      return newItem
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to create item'
      setError(errorMessage)
      console.error('Error creating item:', err)
      return null
    }
  }

  // Update item
  const updateItem = async (guid: string, updateData: ItemUpdateDTO): Promise<Item | null> => {
    try {
      setError(null)

      const updatedItem = await request<Item>(`/items/${guid}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(updateData),
      })

      items.value = items.value.map(item =>
        item.guid === guid ? updatedItem : item
      )

      return updatedItem
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to update item'
      setError(errorMessage)
      console.error('Error updating item:', err)
      return null
    }
  }

  // Delete item
  const deleteItem = async (guid: string): Promise<boolean> => {
    try {
      setError(null)

      await request(`/items/${guid}`, {
        method: 'DELETE',
      })

      items.value = items.value.filter(item => item.guid !== guid)
      return true
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to delete item'
      setError(errorMessage)
      console.error('Error deleting item:', err)
      return false
    }
  }

  return {
    // State
    items,
    loading,
    error,
    searchQuery,

    // Actions
    fetchItems,
    createItem,
    updateItem,
    deleteItem,
    setSearchQuery,
  }
})
