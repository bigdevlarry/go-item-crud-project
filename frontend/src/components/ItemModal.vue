<template>
  <!-- Modal backdrop -->
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
    <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">

      <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="closeModal"></div>

      <!-- Modal panel -->
      <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
        <form @submit.prevent="handleSubmit">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mt-3 text-center sm:mt-0 sm:text-left w-full">
                <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4" id="modal-title">
                  {{ isEditMode ? 'Edit Item' : 'Add New Item' }}
                </h3>

                <!-- Amount and Type Row -->
                <div class="grid grid-cols-12 gap-4 mb-4">
                  <div class="col-span-6">
                    <label for="amount" class="block text-sm font-medium text-gray-700">Amount</label>
                    <input
                      type="number"
                      id="amount"
                      v-model.number="formData.amount"
                      :class="[
                        'mt-1 block w-full rounded-md shadow-sm sm:text-sm',
                        formData.amount <= 0 ? 'border-red-300 focus:ring-red-500 focus:border-red-500' : 'border-gray-300 focus:ring-indigo-500 focus:border-indigo-500'
                      ]"
                      placeholder="Enter amount e.g 100"
                      required
                    />
                    <p v-if="formData.amount <= 0" class="mt-1 text-sm text-red-600">Amount must be greater than 0</p>
                  </div>
                  <div class="col-span-6">
                    <label for="type" class="block text-sm font-medium text-gray-700">Type</label>
                    <select
                      id="type"
                      v-model="formData.type"
                      class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                      required
                    >
                      <option value="ADMISSION">Admission</option>
                      <option value="SUBMISSION">Submission</option>
                      <option value="REVERSAL">Reversal</option>
                    </select>
                  </div>
                </div>

                <!-- Status and Created Date Row -->
                <div class="grid grid-cols-12 gap-4 mb-4">
                  <div class="col-span-6">
                    <label for="status" class="block text-sm font-medium text-gray-700">Status</label>
                    <select
                      id="status"
                      v-model="formData.status"
                      class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                      required
                    >
                      <option value="ACCEPTED">Accepted</option>
                      <option value="DECLINED">Declined</option>
                    </select>
                  </div>
                  <div class="col-span-6">
                    <label for="created" class="block text-sm font-medium text-gray-700">Created (ISO)</label>
                    <input
                      type="datetime-local"
                      id="created"
                      v-model="formData.created"
                      class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                  </div>
                </div>

                <!-- Debtor Information -->
                <div class="mb-4">
                  <h4 class="text-md font-medium text-gray-900 mb-2">Debtor Information</h4>
                  <hr class="border-gray-300 mb-4">
                  <div class="grid grid-cols-12 gap-4 mb-2">
                    <div class="col-span-6">
                      <label for="debtor_first_name" class="block text-sm font-medium text-gray-700">Debtor First Name</label>
                      <input
                        type="text"
                        id="debtor_first_name"
                        v-model="formData.attributes.debtor.first_name"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="First name"
                        required
                      />
                    </div>
                    <div class="col-span-6">
                      <label for="debtor_last_name" class="block text-sm font-medium text-gray-700">Debtor Last Name</label>
                      <input
                        type="text"
                        id="debtor_last_name"
                        v-model="formData.attributes.debtor.last_name"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="Last name"
                        required
                      />
                    </div>
                  </div>
                  <div class="grid grid-cols-12 gap-4">
                    <div class="col-span-6">
                      <label for="debtor_sort_code" class="block text-sm font-medium text-gray-700">Debtor Sort Code</label>
                      <input
                        type="text"
                        id="debtor_sort_code"
                        v-model="debtorSortCode"
                        :class="[
                          'mt-1 block w-full rounded-md shadow-sm sm:text-sm',
                          !validateSortCode(formData.attributes.debtor.account.sort_code) && formData.attributes.debtor.account.sort_code !== ''
                            ? 'border-red-300 focus:ring-red-500 focus:border-red-500'
                            : 'border-gray-300 focus:ring-indigo-500 focus:border-indigo-500'
                        ]"
                        placeholder="00-00-00"
                        maxlength="8"
                        required
                      />
                      <p v-if="!validateSortCode(formData.attributes.debtor.account.sort_code) && formData.attributes.debtor.account.sort_code !== ''" class="mt-1 text-sm text-red-600">Sort code must be in format 00-00-00</p>
                    </div>
                    <div class="col-span-6">
                      <label for="debtor_account_number" class="block text-sm font-medium text-gray-700">Debtor Account Number</label>
                      <input
                        type="text"
                        id="debtor_account_number"
                        v-model="debtorAccountNumber"
                        :class="[
                          'mt-1 block w-full rounded-md shadow-sm sm:text-sm',
                          !validateAccountNumber(formData.attributes.debtor.account.account_number) && formData.attributes.debtor.account.account_number !== ''
                            ? 'border-red-300 focus:ring-red-500 focus:border-red-500'
                            : 'border-gray-300 focus:ring-indigo-500 focus:border-indigo-500'
                        ]"
                        placeholder="00000000"
                        maxlength="8"
                        required
                      />
                      <p v-if="!validateAccountNumber(formData.attributes.debtor.account.account_number) && formData.attributes.debtor.account.account_number !== ''" class="mt-1 text-sm text-red-600">Account number must be 8 digits</p>
                    </div>
                  </div>
                </div>

                <!-- Beneficiary Information -->
                <div class="mb-4">
                  <h4 class="text-md font-medium text-gray-900 mb-2">Beneficiary Information</h4>
                  <hr class="border-gray-300 mb-4">
                  <div class="grid grid-cols-12 gap-4 mb-2">
                    <div class="col-span-6">
                      <label for="beneficiary_first_name" class="block text-sm font-medium text-gray-700">Beneficiary First Name</label>
                      <input
                        type="text"
                        id="beneficiary_first_name"
                        v-model="formData.attributes.beneficiary.first_name"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="First name"
                        required
                      />
                    </div>
                    <div class="col-span-6">
                      <label for="beneficiary_last_name" class="block text-sm font-medium text-gray-700">Beneficiary Last Name</label>
                      <input
                        type="text"
                        id="beneficiary_last_name"
                        v-model="formData.attributes.beneficiary.last_name"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="Last name"
                        required
                      />
                    </div>
                  </div>
                  <div class="grid grid-cols-12 gap-4">
                    <div class="col-span-6">
                      <label for="beneficiary_sort_code" class="block text-sm font-medium text-gray-700">Beneficiary Sort Code</label>
                      <input
                        type="text"
                        id="beneficiary_sort_code"
                        v-model="beneficiarySortCode"
                        :class="[
                          'mt-1 block w-full rounded-md shadow-sm sm:text-sm',
                          !validateSortCode(formData.attributes.beneficiary.account.sort_code) && formData.attributes.beneficiary.account.sort_code !== ''
                            ? 'border-red-300 focus:ring-red-500 focus:border-red-500'
                            : 'border-gray-300 focus:ring-indigo-500 focus:border-indigo-500'
                        ]"
                        placeholder="00-00-00"
                        maxlength="8"
                        required
                      />
                      <p v-if="!validateSortCode(formData.attributes.beneficiary.account.sort_code) && formData.attributes.beneficiary.account.sort_code !== ''" class="mt-1 text-sm text-red-600">Sort code must be in format 00-00-00</p>
                    </div>
                    <div class="col-span-6">
                      <label for="beneficiary_account_number" class="block text-sm font-medium text-gray-700">Beneficiary Account Number</label>
                      <input
                        type="text"
                        id="beneficiary_account_number"
                        v-model="beneficiaryAccountNumber"
                        :class="[
                          'mt-1 block w-full rounded-md shadow-sm sm:text-sm',
                          !validateAccountNumber(formData.attributes.beneficiary.account.account_number) && formData.attributes.beneficiary.account.account_number !== ''
                            ? 'border-red-300 focus:ring-red-500 focus:border-red-500'
                            : 'border-gray-300 focus:ring-indigo-500 focus:border-indigo-500'
                        ]"
                        placeholder="00000000"
                        maxlength="8"
                        required
                      />
                      <p v-if="!validateAccountNumber(formData.attributes.beneficiary.account.account_number) && formData.attributes.beneficiary.account.account_number !== ''" class="mt-1 text-sm text-red-600">Account number must be 8 digits</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Modal footer -->
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              type="submit"
              :disabled="isSubmitting || !isFormValid"
              class="w-full inline-flex justify-center items-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:ml-3 sm:w-auto sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <!-- Loading Spinner -->
              <svg
                v-if="isSubmitting"
                class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                ></circle>
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              {{ isSubmitting ? (isEditMode ? 'Updating...' : 'Creating...') : (isEditMode ? 'Update Item' : 'Create Item') }}
            </button>
            <button
              type="button"
              @click="closeModal"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch } from 'vue'
import { useToast } from 'vue-toastification'
import { useItemsStore } from '../stores/items'
import { formatSortCode, formatAccountNumber, validateSortCode, validateAccountNumber } from '../utils/formatters'
import type { Item, ItemCreateDTO, ItemUpdateDTO } from '@/types'

interface Props {
  isOpen: boolean
  item?: Item | null
}

interface Emits {
  (e: 'close'): void
  (e: 'item-created'): void
  (e: 'item-updated'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const itemsStore = useItemsStore()
const toast = useToast()

const isSubmitting = ref(false)

const formData = reactive<ItemCreateDTO>({
  amount: 0,
  type: 'ADMISSION',
  status: 'ACCEPTED',
  created: new Date().toISOString().slice(0, 16),
  attributes: {
    debtor: {
      first_name: '',
      last_name: '',
      account: {
        sort_code: '',
        account_number: ''
      }
    },
    beneficiary: {
      first_name: '',
      last_name: '',
      account: {
        sort_code: '',
        account_number: ''
      }
    }
  }
})

const isEditMode = computed(() => !!props.item)

// Populate form when item changes (for edit mode)
watch(() => props.item, (newItem) => {
  if (newItem) {
    Object.assign(formData, {
      amount: newItem.amount,
      type: newItem.type,
      status: newItem.status,
      created: new Date(newItem.created).toISOString().slice(0, 16)
    })

    // Deep merge for nested objects
    Object.assign(formData.attributes.debtor, newItem.attributes.debtor)
    Object.assign(formData.attributes.beneficiary, newItem.attributes.beneficiary)
  }
}, { immediate: true })

// Form validation using functional approach
const isFormValid = computed(() => {
  const validations = [
    formData.amount > 0,
    formData.type,
    formData.status,

    // Debtor validations
    formData.attributes.debtor.first_name.trim() !== '',
    formData.attributes.debtor.last_name.trim() !== '',
    validateSortCode(formData.attributes.debtor.account.sort_code),
    validateAccountNumber(formData.attributes.debtor.account.account_number),

    // Beneficiary validations
    formData.attributes.beneficiary.first_name.trim() !== '',
    formData.attributes.beneficiary.last_name.trim() !== '',
    validateSortCode(formData.attributes.beneficiary.account.sort_code),
    validateAccountNumber(formData.attributes.beneficiary.account.account_number)
  ]

  return validations.every(Boolean)
})

const closeModal = () => {
  // Reset form when closing modal
  resetForm()
  emit('close')
}

const resetForm = () => {
  // Reset top-level properties
  Object.assign(formData, {
    amount: 0,
    type: 'ADMISSION',
    status: 'ACCEPTED',
    created: new Date().toISOString().slice(0, 16)
  })

  Object.assign(formData.attributes.debtor, {
    first_name: '',
    last_name: '',
    account: {
      sort_code: '',
      account_number: ''
    }
  })

  Object.assign(formData.attributes.beneficiary, {
    first_name: '',
    last_name: '',
    account: {
      sort_code: '',
      account_number: ''
    }
  })
}

// Handle form submission
const handleSubmit = async () => {
  isSubmitting.value = true

  try {
    // Validate form data before sending
    if (!formData.attributes || !formData.attributes.debtor || !formData.attributes.beneficiary) {
      console.error('Invalid form data structure:', formData)
      toast.error('Form data is invalid. Please try again.')
      return
    }

    if (isEditMode.value && props.item?.guid) {
      // Update existing item
      const updateData: ItemUpdateDTO = {
        amount: formData.amount,
        type: formData.type,
        status: formData.status,
        created: formData.created,
        attributes: formData.attributes
      }
      const updatedItem = await itemsStore.updateItem(props.item.guid, updateData)

      if (updatedItem) {
        toast.success('Item updated successfully!')
        closeModal()
        emit('item-updated')
      } else {
        toast.error('Failed to update item. Please try again.')
      }
    } else {
      // Create new item
      const newItem = await itemsStore.createItem(formData)

      if (newItem) {
        toast.success('Item created successfully!')
        closeModal()
        emit('item-created')
      } else {
        toast.error('Failed to create item. Please try again.')
      }
    }
  } catch (error) {
    console.error('Error processing item:', error)
    toast.error(`Failed to ${isEditMode.value ? 'update' : 'create'} item. Please try again.`)
  } finally {
    isSubmitting.value = false
  }
}

const debtorSortCode = computed({
  get: () => formData.attributes.debtor.account.sort_code,
  set: (value: string) => {
    formData.attributes.debtor.account.sort_code = formatSortCode(value)
  }
})

const debtorAccountNumber = computed({
  get: () => formData.attributes.debtor.account.account_number,
  set: (value: string) => {
    formData.attributes.debtor.account.account_number = formatAccountNumber(value)
  }
})

const beneficiarySortCode = computed({
  get: () => formData.attributes.beneficiary.account.sort_code,
  set: (value: string) => {
    formData.attributes.beneficiary.account.sort_code = formatSortCode(value)
  }
})

const beneficiaryAccountNumber = computed({
  get: () => formData.attributes.beneficiary.account.account_number,
  set: (value: string) => {
    formData.attributes.beneficiary.account.account_number = formatAccountNumber(value)
  }
})
</script>
