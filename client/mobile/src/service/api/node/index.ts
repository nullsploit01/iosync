import { httpClient } from '..'

import { IAPIResponse } from '@/src/types/api'
import { INode } from '@/src/types/models/node'

export const nodeAPIService = {
  getNodes: async () => {
    return await httpClient.get<IAPIResponse<INode[]>>('/nodes')
  }
}
