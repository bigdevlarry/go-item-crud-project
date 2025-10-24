import type { ItemType, ItemStatus } from '@/types'

/**
 * Format a number as currency with a pound symbol
 */
export const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('en-GB', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(amount)
}

/**
 * Format a date string to a readable format
 */
export const formatDate = (dateString: string): string => {
  return new Date(dateString).toLocaleDateString('en-GB', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

/**
 * Get CSS classes for status badge styling
 */
export const getStatusBadgeClass = (status: ItemStatus): string => {
  const baseClasses = 'inline-flex px-2 py-1 text-xs font-semibold rounded-full'

  switch (status) {
    case 'ACCEPTED':
      return `${baseClasses} bg-green-100 text-green-800`
    case 'DECLINED':
      return `${baseClasses} bg-red-100 text-red-800`
    default:
      return `${baseClasses} bg-gray-100 text-gray-800`
  }
}

/**
 * Get CSS classes for type badge styling
 */
export const getTypeBadgeClass = (type: ItemType): string => {
  const baseClasses = 'inline-flex px-2 py-1 text-xs font-semibold rounded-full'

  switch (type) {
    case 'ADMISSION':
      return `${baseClasses} bg-blue-100 text-blue-800`
    case 'SUBMISSION':
      return `${baseClasses} bg-yellow-100 text-yellow-800`
    case 'REVERSAL':
      return `${baseClasses} bg-purple-100 text-purple-800`
    default:
      return `${baseClasses} bg-gray-100 text-gray-800`
  }
}

/**
 * Format sort code to 00-00-00 format
 */
export const formatSortCode = (value: string): string => {
  // Remove all non-numeric characters
  const numbers = value.replace(/\D/g, '')

  // Limit to 6 digits
  const limited = numbers.slice(0, 6)

  // Format as 00-00-00
  if (limited.length <= 2) {
    return limited
  } else if (limited.length <= 4) {
    return `${limited.slice(0, 2)}-${limited.slice(2)}`
  } else {
    return `${limited.slice(0, 2)}-${limited.slice(2, 4)}-${limited.slice(4)}`
  }
}

/**
 * Validate sort code format (00-00-00)
 */
export const validateSortCode = (sortCode: string): boolean => {
  const sortCodeRegex = /^\d{2}-\d{2}-\d{2}$/
  return sortCodeRegex.test(sortCode)
}

/**
 * Format account number and limit to 8 characters
 */
export const formatAccountNumber = (value: string): string => {
  // Remove all non-numeric characters
  const numbers = value.replace(/\D/g, '')

  // Limit to 8 digits
  return numbers.slice(0, 8)
}

/**
 * Validate account number (exactly 8 digits)
 */
export const validateAccountNumber = (accountNumber: string): boolean => {
  const accountRegex = /^\d{8}$/
  return accountRegex.test(accountNumber)
}
