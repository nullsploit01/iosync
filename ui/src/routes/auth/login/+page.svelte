<script lang="ts">
  import { goto } from '$app/navigation'
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import { authService } from '@/services/api/auth'
  import { writable } from 'svelte/store'

  let username = writable('')
  let password = writable('')

  const handleSubmit = async () => {
    let usernameValue, passwordValue

    username.subscribe((value) => (usernameValue = value))()
    password.subscribe((value) => (passwordValue = value))()

    await authService.login(usernameValue!, passwordValue!)
    goto('/')
  }
</script>

<div class="flex h-full place-items-center justify-center">
  <div class="flex flex-col place-items-center justify-center">
    <Input bind:value={$username} placeholder="Username" class="m-3 border-gray-400" />
    <Input
      bind:value={$password}
      type="password"
      placeholder="Password"
      class="m-3 border-gray-400"
    />
    <Button size="lg" class="m-3" on:click={handleSubmit}>Login</Button>
  </div>
</div>
