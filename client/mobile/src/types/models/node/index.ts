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
