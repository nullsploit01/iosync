<script lang="ts">
  import { onMount, setContext } from 'svelte'
  import '../app.css'
  import { ContextKeys } from '@/config/context'
  import { authService } from '@/services/api/auth'
  import Navbar from '@/components/ui/navbar/navbar.svelte'
  import type { ISession } from '@/types/models'
  import { readable } from 'svelte/store'

  const session = readable<ISession | undefined>({} as ISession, (set) => {
    authService
      .me()
      .then(({ data }) => {
        set(data.data)
      })
      .catch((e) => set(undefined))
  })

  setContext(ContextKeys.user.session, {
    ...session
  })
</script>

<div class="h-full bg-gray-200">
  <Navbar />
  <slot></slot>
</div>
