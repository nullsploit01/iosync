<script>
  import { onMount } from 'svelte'
  import { authService } from '@/services/api/auth'
  import { AxiosError } from 'axios'
  import { goto } from '$app/navigation'

  onMount(async () => {
    try {
      await authService.me()
      goto('/')
    } catch (error) {}
  })
</script>

<div class="h-full bg-gray-200 p-2">
  <slot></slot>
</div>
