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

