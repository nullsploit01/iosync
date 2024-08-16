import { goto } from '$app/navigation'
import axios, { AxiosError, type AxiosResponse, HttpStatusCode } from 'axios'

import { API_BASE_URL } from '@/config/environment'

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
