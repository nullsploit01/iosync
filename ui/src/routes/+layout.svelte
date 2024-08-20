<script lang="ts">
  import { onMount, setContext } from 'svelte'
  import { ContextKeys } from '@/config/context'
  import { authService } from '@/services/api/auth'
  import Navbar from '@/components/ui/navbar/navbar.svelte'
  import type { ISession } from '@/types/models'
  import { writable } from 'svelte/store'
  import '../app.css'

  const session = writable<ISession | undefined>(undefined)

  onMount(async () => {
    try {
      const { data } = await authService.me()
      session.set(data.data)
    } catch {
      session.set(undefined)
    }
  })

  setContext(ContextKeys.user.session, session)
</script>

<div class="h-full bg-gray-200">
  <Navbar />
  <slot></slot>
</div>
