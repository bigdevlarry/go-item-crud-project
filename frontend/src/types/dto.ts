import type { ItemType, ItemStatus } from './enums'
import type { Attributes } from './entities'

export interface ItemCreateDTO {
  amount: number
  type: ItemType
  status: ItemStatus
  created?: string
  attributes: Attributes
}

export interface ItemUpdateDTO {
  amount?: number
  type?: ItemType
  status?: ItemStatus
  created?: string
  attributes?: Partial<Attributes>
}

export interface ItemFormData {
  amount: string
  type: string
  status: string
  created: string
  attributes: {
    debtor: {
      first_name: string
      last_name: string
      account: {
        sort_code: string
        account_number: string
      }
    }
    beneficiary: {
      first_name: string
      last_name: string
      account: {
        sort_code: string
        account_number: string
      }
    }
  }
}
