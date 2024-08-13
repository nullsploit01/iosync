<script lang="ts">
  import { goto } from '$app/navigation'
  import { Button } from '@/components/ui/button'
  import { Checkbox } from '@/components/ui/checkbox'
  import { Input } from '@/components/ui/input'
  import { Label } from '@/components/ui/label'
  import { authService } from '@/services/api/auth'
  import { writable } from 'svelte/store'

  let username = writable('')
  let password = writable('')
  let checked: boolean

  const handleSubmit = async () => {
    let usernameValue, passwordValue

    username.subscribe((value) => (usernameValue = value))()
    password.subscribe((value) => (passwordValue = value))()

    await authService.login(usernameValue!, passwordValue!, checked)
    goto('/')
  }
</script>

<div class="flex h-full place-items-center justify-center">
  <div
    class="o-5 m-5 flex w-4/12 flex-col place-items-center justify-center rounded-lg border-2 border-double border-gray-800 p-10"
  >
    <p class="mb-3 text-2xl">Login</p>
    <Input bind:value={$username} placeholder="Username" class="m-3 border-gray-400" />
    <Input
      bind:value={$password}
      type="password"
      placeholder="Password"
      class="m-3 border-gray-400"
    />
    <div class="flex items-center space-x-2">
      <Checkbox id="terms" bind:checked aria-labelledby="terms-label" />
      <Label
        id="terms-label"
        for="terms"
        class="m-1 w-full font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
      >
        Remember Me
      </Label>
    </div>
    <Button size="lg" class="m-3 w-full" on:click={handleSubmit}>Login</Button>
    <Button size="lg" class=" w-full border-gray-900" variant="outline" on:click={() => {}}
      >Register</Button
    >
  </div>
</div>
