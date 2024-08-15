<script lang="ts">
  import { onMount, setContext } from 'svelte'
  import '../app.css'
  import { ContextKeys } from '@/config/context'
  import { authService } from '@/services/api/auth'
  import Navbar from '@/components/ui/navbar/navbar.svelte'

  let session: any

  onMount(async () => {
    const { data } = await authService.me()
    session = data
  })

  setContext(ContextKeys.user.session, {
    session
  })
</script>

<div class="h-full bg-gray-200">
  <Navbar />
  <slot></slot>
</div>
