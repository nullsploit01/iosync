export interface IUser {
  id: number
  username: string
  name: string
  is_active: boolean
  last_login: Date
  created_at: Date
  updated_at: Date
}

export interface IDevice {
  id: string
  username: string
  name: string
  created_at: string
  updated_at: string
}

export interface ISession {
  sessionId: string
  username: string
  expiresAt: string
}
