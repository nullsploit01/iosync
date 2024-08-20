<script lang="ts">
  import { getContext, onDestroy } from 'svelte'
  import { goto } from '$app/navigation'
  import { ContextKeys } from '@/config/context'
  import type { ISession } from '@/types/models'
  import type { Writable } from 'svelte/store'

  const session = getContext<Writable<ISession | undefined>>(ContextKeys.user.session)

  const unsubscribe = session.subscribe((value) => {
    if (value) {
      goto('/')
    }
  })

  onDestroy(unsubscribe)
</script>

<div>
  <slot></slot>
</div>
