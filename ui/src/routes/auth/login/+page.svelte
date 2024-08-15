<script lang="ts">
  import { goto } from '$app/navigation'
  import { Button } from '@/components/ui/button'
  import { Checkbox } from '@/components/ui/checkbox'
  import { Input } from '@/components/ui/input'
  import { Label } from '@/components/ui/label'
  import { ContextKeys } from '@/config/context'
  import { authService } from '@/services/api/auth'
  import { onMount, getContext } from 'svelte'

  let username = ''
  let password = ''
  let checked: boolean

  const handleSubmit = async () => {
    await authService.login(username, password, checked)
    goto('/')
  }
</script>

<div class="mt-12 flex h-full place-items-center justify-center">
  <div
    class="max-w-2/12 m-5 flex flex-col place-items-center justify-center rounded-lg border-2 border-double border-gray-800 p-10"
  >
    <p class="mb-3 text-2xl">Login</p>
    <Input bind:value={username} placeholder="Username" class="m-3 border-gray-400" />
    <Input
      bind:value={password}
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
    <p class="font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
      Dont have an account? <a class="text-blue-500" href="/auth/register" on:click={() => {}}
        >Sign Up</a
      >
    </p>
  </div>
</div>
