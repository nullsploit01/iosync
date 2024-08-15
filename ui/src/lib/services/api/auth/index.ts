import { httpClient } from '@/clients/http'
import type { IApiRespose } from '@/types/api'
import type { ISession } from '@/types/models'

export const authService = {
  me: async () => {
    return await httpClient.get<IApiRespose<ISession>>('/auth/me')
  },

  login: async (username: string, password: string, rememberMe: boolean) => {
    return await httpClient.post<IApiRespose<ISession>>('/auth/login', {
      username,
      password,
      rememberMe
    })
  },

  register: async (name: string, username: string, password: string) => {
    return await httpClient.post<IApiRespose<ISession>>('/auth/register', {
      name,
      username,
      password
    })
  }
}
