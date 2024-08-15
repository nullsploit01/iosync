import { httpClient } from '@/clients/http'
import type { IApiRespose } from '@/types/api'
import type { IDevice } from '@/types/models'

export const deviceService = {
  getDevices: async () => {
    return await httpClient.get<IApiRespose<IDevice[]>>('/devices')
  }
}
