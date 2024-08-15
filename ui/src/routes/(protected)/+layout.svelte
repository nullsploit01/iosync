<script lang="ts">
  import { getContext, onDestroy, onMount } from 'svelte'
  import { ContextKeys } from '@/config/context'
  import { goto } from '$app/navigation'
  import type { ISession } from '@/types/models'
  import type { Readable } from 'svelte/store'

  const unsubscribe = getContext<Readable<ISession | undefined>>(
    ContextKeys.user.session
  ).subscribe((value) => {
    if (!value) goto('/auth/login')
  })

  onDestroy(unsubscribe)
</script>

<div class="bg-gray-200 p-2">
  <slot></slot>
</div>
