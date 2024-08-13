import { httpClient } from '@/clients/http'

export const authService = {
  me: async () => {
    return await httpClient.get('/auth/me')
  },

  login: async (username: string, password: string, rememberMe: boolean) => {
    return await httpClient.post('/auth/login', { username, password, rememberMe })
  },

  register: async (name: string, username: string, password: string) => {
    return await httpClient.post('/auth/register', { name, username, password })
  }
}
