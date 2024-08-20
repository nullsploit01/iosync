<script lang="ts">
  import { getContext, onDestroy } from 'svelte'
  import { goto } from '$app/navigation'
  import { ContextKeys } from '@/config/context'
  import type { ISession } from '@/types/models'
  import type { Writable } from 'svelte/store'

  const session = getContext<Writable<ISession | undefined>>(ContextKeys.user.session)

  const unsubscribe = session.subscribe((value) => {
    if (!value) {
      goto('/auth/login')
    }
  })

  onDestroy(unsubscribe)
</script>

<div class="bg-gray-200 p-2">
  <slot></slot>
</div>
