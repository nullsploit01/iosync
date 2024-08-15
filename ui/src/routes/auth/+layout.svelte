<script>
  import { getContext, onMount } from 'svelte'
  import { authService } from '@/services/api/auth'
  import { AxiosError } from 'axios'
  import { goto } from '$app/navigation'
  import { ContextKeys } from '@/config/context'

  onMount(async () => {
    try {
      await authService.me()
      goto('/')
    } catch (error) {}
  })
</script>

<div>
  <slot></slot>
</div>
