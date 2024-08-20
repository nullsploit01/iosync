import type { Load } from '@sveltejs/kit'

import { authService } from '@/services/api/auth'

export const load: Load = async ({ fetch }) => {
  try {
    const { data } = await authService.me()
    if (!data.data) {
      return { status: 302, redirect: '/auth/login' }
    }
    return { props: { session: data.data } }
  } catch {
    return { status: 302, redirect: '/auth/login' }
  }
}
