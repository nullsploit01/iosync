<script>
  import { onMount } from 'svelte'
  import { authService } from '@/services/api/auth'

  onMount(async () => {
    await authService.me()
  })
</script>

<div class="h-full bg-gray-200 p-2">
  <slot></slot>
</div>
