import { httpClient } from '..'

export const nodeAPIService = {
  getNodes: async () => {
    return await httpClient.get('/nodes')
  }
}
