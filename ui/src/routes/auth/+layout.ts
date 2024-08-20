import type { Load } from '@sveltejs/kit'

import { authService } from '@/services/api/auth'

export const load: Load = async ({ fetch }) => {
  try {
    const { data } = await authService.me()
    if (data.data) {
      return { status: 302, redirect: '/' }
    }
    return {}
  } catch {
    return {}
  }
}
