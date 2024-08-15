export interface IApiRespose<T> {
  error: boolean
  message?: string
  data?: T
}
