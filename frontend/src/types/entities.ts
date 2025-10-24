import type { ItemType, ItemStatus } from './enums'

export interface Item {
  guid: string
  index: number
  amount: number
  type: ItemType
  status: ItemStatus
  created: string
  attributes: Attributes
}

export interface Attributes {
  debtor: Party
  beneficiary: Party
}

export interface Party {
  first_name: string
  last_name: string
  account: Account
}

export interface Account {
  sort_code: string
  account_number: string
}
