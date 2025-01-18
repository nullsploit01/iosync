import { httpClient } from '..'

import { IAPIResponse } from '@/src/types/api'
import { INode, INodeAPIValue, INodeWithAPIKeys } from '@/src/types/models/node'

export const nodeAPIService = {
  getNodes: async () => {
    return await httpClient.get<IAPIResponse<INode[]>>('/nodes')
  },

  getNode: async (id: number) => {
    return await httpClient.get<IAPIResponse<INodeWithAPIKeys>>(`/nodes/${id}`)
  },

  getNodeValuesByApiKey: async (apiKey: string) => {
    return await httpClient.get<IAPIResponse<INodeAPIValue[]>>(`/nodes/${apiKey}/values`)
  }
}
