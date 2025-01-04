export interface IAPIResponse<T> {
  error: boolean
  message: string
  data: T
  pagination: {
    current_page: number
    total_pages: number
    per_page: number
    total_items: number
  }
}
