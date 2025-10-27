const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

// Request helper
export async function request<T>(endpoint: string, options?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${endpoint}`, options)

  if (!response.ok) {
    let msg = `HTTP error! status: ${response.status}`
    const err = await response.json().catch(() => ({}))
    if (err.error) msg = err.error
    throw new Error(msg)
  }

  if (response.status === 204 || response.headers.get('content-length') === '0') {
    return {} as T
  }

  return await response.json() as Promise<T>
}

