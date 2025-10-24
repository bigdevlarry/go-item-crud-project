<template>
  <div class="bg-white shadow-sm ring-1 ring-gray-900/5 sm:rounded-xl">
            <!-- Header -->
            <div class="px-4 py-6 sm:px-6">
              <div class="flex items-center justify-between">
                <div>
                  <h1 class="text-2xl font-semibold text-gray-900">Items</h1>
                  <p class="mt-1 text-sm text-gray-500">
                    A list of all the items in your account including their details, type, and status.
                  </p>
                </div>
                <div class="flex items-center space-x-4">
                  <!-- Search Component -->
                  <ItemSearch v-model="searchQuery" />
                  <button
                    @click="openCreateModal"
                    class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  >
                    Add item
                  </button>
                </div>
              </div>
            </div>

    <!-- Table -->
    <div class="overflow-hidden">
      <table class="min-w-full divide-y divide-gray-300">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              #
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Amount
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Type
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Status
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Debtor
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Beneficiary
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Created
            </th>
            <th scope="col" class="relative px-6 py-3">
              <span class="sr-only">Actions</span>
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading" class="animate-pulse">
            <td colspan="8" class="px-6 py-4 text-center text-gray-500">
              Loading items...
            </td>
          </tr>
          <tr v-else-if="items.length === 0">
            <td colspan="8" class="px-6 py-4 text-center text-gray-500">
              {{ searchQuery ? 'No items match your search.' : 'No items found. Create your first item!' }}
            </td>
          </tr>
                  <tr v-else v-for="(item, index) in items" :key="item.guid" class="hover:bg-gray-50">
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      #{{ index + 1 }}
                    </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              ${{ item.amount }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full"
                    :class="getTypeBadgeClass(item.type)">
                {{ item.type }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full"
                    :class="getStatusBadgeClass(item.status)">
                {{ item.status }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ item.attributes.debtor.first_name }} {{ item.attributes.debtor.last_name }}
              <p>{{ item.attributes.debtor.account.sort_code }} . {{ item.attributes.debtor.account.account_number }}</p>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ item.attributes.beneficiary.first_name }} {{ item.attributes.beneficiary.last_name }}
              <p> {{ item.attributes.beneficiary.account.sort_code }} . {{ item.attributes.beneficiary.account.account_number }}</p>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ formatDate(item.created) }}
            </td>
                    <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                      <button
                        @click="editItem(item)"
                        class="text-indigo-600 hover:text-indigo-900 mr-3 p-1 rounded-md hover:bg-indigo-50"
                        title="Edit item"
                      >
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                        </svg>
                      </button>
                      <button
                        @click="deleteItem(item.guid)"
                        class="text-red-600 hover:text-red-900 p-1 rounded-md hover:bg-red-50"
                        title="Delete item"
                      >
                        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                        </svg>
                      </button>
                    </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Add Item Modal -->
    <AddItemModal
      :is-open="showCreateModal"
      @close="closeCreateModal"
      @item-created="handleItemCreated"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import AddItemModal from './AddItemModal.vue'
import ItemSearch from './ItemSearch.vue'
import type { Item } from '../types/entities'
import type { ItemType, ItemStatus } from '../types/enums'
import type { BadgeStyle } from '../types/components'

        const items = ref<Item[]>([])
        const loading = ref(false)
        const showCreateModal = ref(false)
        const searchQuery = ref('')

        // Fetch items from API with search query
const fetchItems = async () => {
  loading.value = true
  try {
    const url = new URL('http://localhost:8080/items')
    if (searchQuery.value.trim()) {
      url.searchParams.append('query', searchQuery.value.trim())
    }
    
    const response = await fetch(url.toString())
    const data = await response.json()
    items.value = data || []
  } catch (error) {
    console.error('Error fetching items:', error)
  } finally {
    loading.value = false
  }
}

        // Watch for search query changes and refetch items
        watch(searchQuery, () => {
          fetchItems()
        })

        // Badge styling functions
        const getTypeBadgeClass = (type: ItemType): string => {
          const classes: Record<ItemType, string> = {
            'ADMISSION': 'bg-blue-100 text-blue-800',
            'DISCHARGE': 'bg-green-100 text-green-800',
            'TRANSFER': 'bg-yellow-100 text-yellow-800'
          }
          return classes[type] || 'bg-gray-100 text-gray-800'
        }

        const getStatusBadgeClass = (status: ItemStatus): string => {
          const classes: Record<ItemStatus, string> = {
            'PENDING': 'bg-yellow-100 text-yellow-800',
            'ACCEPTED': 'bg-green-100 text-green-800',
            'REJECTED': 'bg-red-100 text-red-800'
          }
          return classes[status] || 'bg-gray-100 text-gray-800'
        }

// Date formatting
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// CRUD operations
const openCreateModal = () => {
  showCreateModal.value = true
}

const closeCreateModal = () => {
  showCreateModal.value = false
}

const handleItemCreated = () => {
  fetchItems() // Refresh the list
}

const editItem = (item: Item) => {
  // TODO: Implement edit modal
  console.log('Edit item:', item)
}

const deleteItem = async (guid: string) => {
  if (confirm('Are you sure you want to delete this item?')) {
    try {
      const response = await fetch(`http://localhost:8080/items/${guid}`, {
        method: 'DELETE'
      })
      if (response.ok) {
        await fetchItems() // Refresh the list
      }
    } catch (error) {
      console.error('Error deleting item:', error)
    }
  }
}

// Load items on component mount
onMounted(() => {
  fetchItems()
})
</script>
