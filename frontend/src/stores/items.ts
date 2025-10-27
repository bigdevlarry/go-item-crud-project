import {defineStore} from 'pinia'
import {ref} from 'vue'
import type {Item, ItemCreateDTO, ItemUpdateDTO} from '@/types'
import {request} from '@/utils/request'

export const useItemsStore = defineStore(
  'items', () => {
  // State
  const items = ref<Item[]>([])
  const loading = ref(false)
  const searchQuery = ref('')

  // Actions
  const setLoading = (isLoading: boolean) => {
    loading.value = isLoading
  }

  const setSearchQuery = (query: string) => {
    searchQuery.value = query
  }

  // Fetch all items
  const fetchItems = async (query?: string) => {
    try {
      setLoading(true)

      items.value = []

      const endpoint = query?.trim() ? `/items?query=${encodeURIComponent(query.trim())}` : '/items'
      items.value = await request<Item[]>(endpoint)
    } catch (err) {
      items.value = []
      throw err
    } finally {
      setLoading(false)
    }
  }

  // Create new item
  const createItem = async (itemData: ItemCreateDTO): Promise<Item | null> => {
    try {
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
      console.error('Error creating item:', err)
      return null
    }
  }

  // Update item
  const updateItem = async (guid: string, updateData: ItemUpdateDTO): Promise<Item | null> => {
    try {
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
      console.error('Error updating item:', err)
      return null
    }
  }

  // Delete item
  const deleteItem = async (guid: string): Promise<boolean> => {
    try {
      await request(`/items/${guid}`, {
        method: 'DELETE',
      })

      items.value = items.value.filter(item => item.guid !== guid)
      return true
    } catch (err) {
      console.error('Error deleting item:', err)
      return false
    }
  }

  return {
    // State
    items,
    loading,
    searchQuery,

    // Actions
    fetchItems,
    createItem,
    updateItem,
    deleteItem,
    setSearchQuery,
  }
})
