import type { Item } from './entities'

export interface ApiResponse<T> {
  message: string
  data?: T
}

export interface ItemsResponse {
  items: Item[]
}

export interface ItemResponse {
  item: Item
}

export interface ApiError {
  message: string
  status: number
  details?: Record<string, any>
}
