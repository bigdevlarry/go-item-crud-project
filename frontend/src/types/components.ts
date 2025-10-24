import type { Item } from './entities'

export interface ItemsTableProps {
  items: Item[]
  loading?: boolean
  onEdit?: (item: Item) => void
  onDelete?: (guid: string) => void
}

export interface AddItemModalProps {
  isOpen: boolean
  onClose: () => void
  onItemCreated: () => void
}

export interface ItemSearchProps {
  modelValue: string
  placeholder?: string
}

export interface TableColumn {
  key: string
  label: string
}

export interface BadgeStyle {
  bg: string
  text: string
}
