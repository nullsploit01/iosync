import { goto } from '$app/navigation'
import { API_BASE_URL } from '@/config/environment'
import axios, { AxiosError, HttpStatusCode, type AxiosResponse } from 'axios'

const httpClient = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true
})

const responseInterceptor = (response: AxiosResponse) => {
  return response
}

const errorInterceptor = (error: AxiosError) => {
  switch (error.response?.status) {
    case HttpStatusCode.Unauthorized: {
      goto('/auth/login')
      break
    }

    case HttpStatusCode.NotFound: {
      break
    }
  }

  return Promise.reject(error)
}

httpClient.interceptors.response.use(responseInterceptor, errorInterceptor)
export { httpClient }
