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
          <svg class="size-6 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v6m3-3H9m12 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
          </svg>

            New item
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
            <td colspan="6" class="px-6 py-4 text-center text-gray-500">
              {{ searchQuery ? 'No items match your search.' : 'No items found. Create your first item!' }}
            </td>
          </tr>
          <template v-else v-for="(item, index) in items" :key="item.guid">
            <!-- Main row -->
            <tr
              @click="toggleItemDetails(item.guid)"
              class="hover:bg-gray-50 cursor-pointer"
              :class="{ 'bg-blue-50': selectedItemGuid === item.guid }"
            >
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
                {{ formatDate(item.created) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium" @click.stop>
                <button
                  @click="editItem(item)"
                  class="text-indigo-600 hover:text-indigo-900 mr-3 p-1 rounded-md hover:bg-indigo-50"
                  title="Edit item"
                >
                  <svg class="size-6 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button
                  @click="deleteItem(item.guid)"
                  class="text-red-600 hover:text-red-900 p-1 rounded-md hover:bg-red-50"
                  title="Delete item"
                >
                  <svg class="size-6 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </td>
            </tr>

            <!-- Expanded details row -->
            <tr v-if="selectedItemGuid === item.guid" class="bg-blue-50">
              <td colspan="6" class="px-6 py-4">
                <div class="space-y-6">
                  <!-- Item Details -->
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <!-- Debtor Details -->
                    <div class="bg-white p-4 rounded-lg shadow-sm">
                      <h4 class="text-sm font-semibold text-gray-900 mb-3">Debtor Details</h4>
                      <div class="space-y-2">
                        <p class="text-sm"><span class="font-medium">Name:</span> {{ item.attributes?.debtor?.first_name || 'N/A' }} {{ item.attributes?.debtor?.last_name || 'N/A' }}</p>
                        <p class="text-sm"><span class="font-medium">Sort Code:</span> {{ item.attributes?.debtor?.account?.sort_code || 'N/A' }}</p>
                        <p class="text-sm"><span class="font-medium">Account Number:</span> {{ item.attributes?.debtor?.account?.account_number || 'N/A' }}</p>
                      </div>
                    </div>

                    <!-- Beneficiary Details -->
                    <div class="bg-white p-4 rounded-lg shadow-sm">
                      <h4 class="text-sm font-semibold text-gray-900 mb-3">Beneficiary Details</h4>
                      <div class="space-y-2">
                        <p class="text-sm"><span class="font-medium">Name:</span> {{ item.attributes?.beneficiary?.first_name || 'N/A' }} {{ item.attributes?.beneficiary?.last_name || 'N/A' }}</p>
                        <p class="text-sm"><span class="font-medium">Sort Code:</span> {{ item.attributes?.beneficiary?.account?.sort_code || 'N/A' }}</p>
                        <p class="text-sm"><span class="font-medium">Account Number:</span> {{ item.attributes?.beneficiary?.account?.account_number || 'N/A' }}</p>
                      </div>
                    </div>
                  </div>

                  <!-- GUID -->
                  <div class="bg-white p-4 rounded-lg shadow-sm">
                    <h4 class="text-sm font-semibold text-gray-900 mb-2">GUID</h4>
                    <p class="text-sm font-mono text-gray-600">{{ item.guid }}</p>
                  </div>

                </div>
              </td>
            </tr>
          </template>
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

// Pinia store
const itemsStore = useItemsStore()

const toast = useToast()

const showCreateModal = ref(false)
const showEditModal = ref(false)
const showDeleteModal = ref(false)
const itemToEdit = ref<Item | null>(null)
const itemToDelete = ref<string | null>(null)
const selectedItemGuid = ref<string | null>(null)

// Computed properties from store
const items = computed(() => itemsStore.items)
const loading = computed(() => itemsStore.loading)
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
  closeModal()
  await itemsStore.fetchItems() // refresh items list
}

const handleItemUpdated = async () => {
  closeModal()
  await itemsStore.fetchItems() // refresh items list
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

// Collapse functionality
const toggleItemDetails = (guid: string) => {
  selectedItemGuid.value = selectedItemGuid.value === guid ? null : guid
}

onMounted(() => {
  itemsStore.fetchItems()
})

</script>
