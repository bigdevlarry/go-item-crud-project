export interface ValidationError {
  field: string
  message: string
  value?: any
}

export interface ValidationResult {
  isValid: boolean
  errors: ValidationError[]
}
