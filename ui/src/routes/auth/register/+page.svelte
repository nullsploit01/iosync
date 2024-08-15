<script lang="ts">
  import { goto } from '$app/navigation'
  import { Button } from '@/components/ui/button'
  import { Input } from '@/components/ui/input'
  import { authService } from '@/services/api/auth'

  let username = ''
  let password = ''
  let name = ''

  const handleSubmit = async () => {
    await authService.register(name, username, password)
    goto('/')
  }
</script>

<div class="mt-12 flex place-items-center justify-center">
  <div
    class="max-w-2/12 m-5 flex flex-col place-items-center justify-center rounded-lg border-2 border-double border-gray-800 p-10"
  >
    <p class="mb-3 text-2xl">Sign Up</p>
    <Input bind:value={name} placeholder="Name" class="m-3 border-gray-400" />
    <Input bind:value={username} placeholder="Username" class="m-3 border-gray-400" />
    <Input
      bind:value={password}
      type="password"
      placeholder="Password"
      class="m-3 border-gray-400"
    />

    <Button size="lg" class="m-3 w-full" on:click={handleSubmit}>Sign Up</Button>
    <p class="font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
      Already have an account? <a class="text-blue-500" href="/auth/login" on:click={() => {}}
        >Login</a
      >
    </p>
  </div>
</div>
