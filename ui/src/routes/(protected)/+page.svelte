<script>
  import { deviceService } from '@/services/api/device'
  import { onMount } from 'svelte'

  const loadDevices = async () => {
    const { data } = await deviceService.getDevices()
  }

  onMount(async () => await loadDevices())
</script>

<h1>Welcome to SvelteKit</h1>
