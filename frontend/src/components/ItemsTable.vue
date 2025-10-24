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
              {{ index + 1 }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              £{{ formatCurrency(item.amount) }}
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
              {{ item.attributes?.debtor?.first_name || 'N/A' }} {{ item.attributes?.debtor?.last_name || 'N/A' }}
              <p>{{ item.attributes?.debtor?.account?.sort_code || 'N/A' }} . {{ item.attributes?.debtor?.account?.account_number || 'N/A' }}</p>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ item.attributes?.beneficiary?.first_name || 'N/A' }} {{ item.attributes?.beneficiary?.last_name || 'N/A' }}
              <p> {{ item.attributes?.beneficiary?.account?.sort_code || 'N/A' }} . {{ item.attributes?.beneficiary?.account?.account_number || 'N/A' }}</p>
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

    <!-- Item Modal (Create/Edit) -->
    <ItemModal
      :is-open="showCreateModal || showEditModal"
      :item="itemToEdit"
      @close="closeModal"
      @item-created="handleItemCreated"
      @item-updated="handleItemUpdated"
    />

    <!-- Delete Confirmation Modal -->
    <ConfirmationModal
      :is-open="showDeleteModal"
      title="Delete Item"
      message="Are you sure you want to delete this item? This action cannot be undone."
      confirm-text="Delete"
      @confirm="confirmDelete"
      @cancel="cancelDelete"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useToast } from 'vue-toastification'
import { useItemsStore } from '../stores/items'
import ItemModal from './ItemModal.vue'
import ItemSearch from './ItemSearch.vue'
import ConfirmationModal from './ConfirmationModal.vue'
import { formatCurrency, formatDate, getStatusBadgeClass, getTypeBadgeClass } from '../utils/formatters'
import type { Item } from '@/types'
import type { ItemType, ItemStatus } from '../types/enums'

// Pinia store
const itemsStore = useItemsStore()

// Toast notifications
const toast = useToast()

// Local state
const showCreateModal = ref(false)
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const itemToEdit = ref<Item | null>(null)
const itemToDelete = ref<string | null>(null)

// Computed properties from store
const items = computed(() => itemsStore.items)
const loading = computed(() => itemsStore.loading)
const error = computed(() => itemsStore.error)
const searchQuery = computed({
  get: () => itemsStore.searchQuery,
  set: (value: string) => itemsStore.setSearchQuery(value)
})

// Watch for search query changes and refetch items
watch(searchQuery, (newQuery) => {
  itemsStore.fetchItems(newQuery)
})


const openCreateModal = () => {
  itemToEdit.value = null
  showCreateModal.value = true
}

const editItem = (item: Item) => {
  itemToEdit.value = item
  showEditModal.value = true
}

const closeModal = () => {
  showCreateModal.value = false
  showEditModal.value = false
  itemToEdit.value = null
}

const handleItemCreated = async () => {
  // Close the modal after successful item creation
  closeModal()
  
  // Refresh the items list to ensure UI updates
  await itemsStore.fetchItems()
}

const handleItemUpdated = async () => {
  // Close the modal after successful item update
  closeModal()
  
  // Refresh the items list to ensure UI updates
  await itemsStore.fetchItems()
}

const deleteItem = (guid: string) => {
  itemToDelete.value = guid
  showDeleteModal.value = true
}

const confirmDelete = async () => {
  if (itemToDelete.value) {
    const success = await itemsStore.deleteItem(itemToDelete.value)
    if (success) {
      toast.success('Item deleted successfully!')
    } else {
      toast.error('Failed to delete item. Please try again.')
    }
    
    cancelDelete()
  }
}

const cancelDelete = () => {
  showDeleteModal.value = false
  itemToDelete.value = null
}


onMounted(() => {
  itemsStore.fetchItems()
})

</script>
