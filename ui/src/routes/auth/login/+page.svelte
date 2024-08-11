<script lang="ts">
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import { API_BASE_URL } from '@/config'
  import { writable } from 'svelte/store'

  let username = writable('')
  let password = writable('')

  async function handleSubmit() {
    let usernameValue, passwordValue

    // Subscribe to the writable stores to get the values
    username.subscribe((value) => (usernameValue = value))()
    password.subscribe((value) => (passwordValue = value))()

    const loginData = {
      username: usernameValue,
      password: passwordValue
    }
    const response = await fetch(`${API_BASE_URL}/api/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(loginData)
    })

    if (!response.ok) {
      throw new Error('Login failed')
    }

    const data = await response.json()
    console.log('Login successful:', data)
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
