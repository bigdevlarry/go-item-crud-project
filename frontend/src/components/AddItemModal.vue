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
                  Add New Item
                </h3>

                <!-- Amount and Type Row -->
                <div class="grid grid-cols-12 gap-4 mb-4">
                  <div class="col-span-6">
                    <label for="amount" class="block text-sm font-medium text-gray-700">Amount</label>
                    <input
                      type="text"
                      id="amount"
                      v-model="formData.amount"
                      class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                      placeholder="Enter amount e.g 100"
                      required
                    />
                  </div>
                  <div class="col-span-6">
                    <label for="type" class="block text-sm font-medium text-gray-700">Type</label>
                    <select
                      id="type"
                      v-model="formData.type"
                      class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                      required
                    >
                      <option value="">Select type</option>
                      <option value="ADMISSION">Admission</option>
                      <option value="DISCHARGE">Discharge</option>
                      <option value="TRANSFER">Transfer</option>
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
                      <option value="">Select status</option>
                      <option value="PENDING">Pending</option>
                      <option value="ACCEPTED">Accepted</option>
                      <option value="REJECTED">Rejected</option>
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
                        v-model="formData.attributes.debtor.account.sort_code"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="00-00-00"
                        required
                      />
                    </div>
                    <div class="col-span-6">
                      <label for="debtor_account_number" class="block text-sm font-medium text-gray-700">Debtor Account Number</label>
                      <input
                        type="text"
                        id="debtor_account_number"
                        v-model="formData.attributes.debtor.account.account_number"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="00000000"
                        required
                      />
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
                        v-model="formData.attributes.beneficiary.account.sort_code"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="00-00-00"
                        required
                      />
                    </div>
                    <div class="col-span-6">
                      <label for="beneficiary_account_number" class="block text-sm font-medium text-gray-700">Beneficiary Account Number</label>
                      <input
                        type="text"
                        id="beneficiary_account_number"
                        v-model="formData.attributes.beneficiary.account.account_number"
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="00000000"
                        required
                      />
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
              :disabled="isSubmitting"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:ml-3 sm:w-auto sm:text-sm disabled:opacity-50"
            >
              {{ isSubmitting ? 'Creating...' : 'Create Item' }}
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
import { ref, reactive } from 'vue'
import type { ItemFormData } from '../types/dto'
import type { AddItemModalProps } from '../types/components'

interface Props {
  isOpen: boolean
}

interface Emits {
  (e: 'close'): void
  (e: 'item-created'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const isSubmitting = ref(false)

// Form data structure
const formData = reactive({
  amount: '',
  type: '',
  status: '',
  created: '',
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

// Close modal
const closeModal = () => {
  emit('close')
  resetForm()
}

// Reset form to initial state
const resetForm = () => {
  formData.amount = ''
  formData.type = ''
  formData.status = ''
  formData.created = ''
  formData.attributes.debtor.first_name = ''
  formData.attributes.debtor.last_name = ''
  formData.attributes.debtor.account.sort_code = ''
  formData.attributes.debtor.account.account_number = ''
  formData.attributes.beneficiary.first_name = ''
  formData.attributes.beneficiary.last_name = ''
  formData.attributes.beneficiary.account.sort_code = ''
  formData.attributes.beneficiary.account.account_number = ''
}

// Handle form submission
const handleSubmit = async () => {
  isSubmitting.value = true

  try {
    const response = await fetch('http://localhost:8080/items', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData)
    })

    if (response.ok) {
      emit('item-created')
      closeModal()
    } else {
      const errorData = await response.json()
      console.error('Error creating item:', errorData)
      alert('Failed to create item. Please try again.')
    }
  } catch (error) {
    console.error('Error creating item:', error)
    alert('Failed to create item. Please try again.')
  } finally {
    isSubmitting.value = false
  }
}
</script>
