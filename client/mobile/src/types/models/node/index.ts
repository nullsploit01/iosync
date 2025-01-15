export interface INode {
  id: number
  name: string
  description: string
  identifier: string
  is_active: boolean
  is_online: boolean
  last_online_at: Date
  created_at: Date
  updated_at: Date
}

export interface INodeAPIKey {
  id: number
  api_key: string
  description: string
  created_at: Date
  updated_at: Date
}

export interface INodeAPIValue {
  id: number
  value: string
  created_at: Date
  updated_at: Date
}

export interface INodeWithAPIKeys extends INode {
  api_keys: INodeAPIKey[]
}
