import type { Item } from './entities'

export interface ItemEvent {
  type: 'create' | 'update' | 'delete'
  item: Item
  timestamp: string
}
