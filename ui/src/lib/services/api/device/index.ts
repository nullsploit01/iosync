import { httpClient } from '@/clients/http'

export const deviceService = {
  getDevices: async () => {
    return await httpClient.get('/devices')
  }
}
