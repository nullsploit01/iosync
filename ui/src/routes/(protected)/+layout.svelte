<script>
  import { getContext } from 'svelte'
  import { ContextKeys } from '@/config/context'
  import { goto } from '$app/navigation'

  const session = getContext(ContextKeys.user.session).session
  if (!session) goto('/auth/login')
</script>

<div class="bg-gray-200 p-2">
  <slot></slot>
</div>
