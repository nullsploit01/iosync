import axios from 'axios'

import { Environment } from '@/src/config/environment'

export const httpClient = axios.create({
  baseURL: Environment.API_URL
})
