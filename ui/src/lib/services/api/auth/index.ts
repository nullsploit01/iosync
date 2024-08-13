import { httpClient } from '@/clients/http'

export const authService = {
  me: async () => {
    return await httpClient.get('/auth/me')
  },

  login: async (username: string, password: string) => {
    return await httpClient.post('/auth/login', { username, password })
  }
}
