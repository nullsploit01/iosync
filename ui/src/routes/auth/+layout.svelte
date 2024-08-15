<script lang="ts">
  import { getContext, onDestroy, onMount } from 'svelte'
  import { authService } from '@/services/api/auth'
  import { AxiosError } from 'axios'
  import { goto } from '$app/navigation'
  import { ContextKeys } from '@/config/context'
  import type { ISession } from '@/types/models'
  import type { Readable } from 'svelte/motion'

  const unsubscribe = getContext<Readable<ISession | undefined>>(
    ContextKeys.user.session
  ).subscribe((value) => {
    if (value) goto('/')
  })

  onDestroy(unsubscribe)
</script>

<div>
  <slot></slot>
</div>
