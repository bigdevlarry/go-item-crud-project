import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import type {Item, ItemCreateDTO, ItemUpdateDTO} from '@/types'

const API_BASE_URL = 'http://localhost:8080'

export const useItemsStore = defineStore('items', () => {
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

      const url = new URL(`${API_BASE_URL}/items`)
      if (query?.trim()) {
        url.searchParams.append('query', query.trim())
      }

      const response = await fetch(url.toString())

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()

      // Handle different response formats
      if (Array.isArray(data)) {
        items.value = data
      } else if (data.items && Array.isArray(data.items)) {
        items.value = data.items
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

      const response = await fetch(`${API_BASE_URL}/items`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(itemData),
      })

      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.error || `HTTP error! status: ${response.status}`)
      }

      const newItem = await response.json()
      items.value.push(newItem)

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

      const response = await fetch(`${API_BASE_URL}/items/${guid}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(updateData),
      })

      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.error || `HTTP error! status: ${response.status}`)
      }

      const updatedItem = await response.json()

      const index = items.value.findIndex(item => item.guid === guid)
      if (index !== -1) {
        items.value[index] = updatedItem
      }

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

      const response = await fetch(`${API_BASE_URL}/items/${guid}`, {
        method: 'DELETE',
      })

      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.error || `HTTP error! status: ${response.status}`)
      }

      const index = items.value.findIndex(item => item.guid === guid)
      if (index !== -1) {
        items.value.splice(index, 1)
      }

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
