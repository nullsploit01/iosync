<script lang="ts">
  import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle
  } from '@/components/ui/card'
  import { deviceService } from '@/services/api/device'
  import type { IDevice } from '@/types/models'
  import { onMount } from 'svelte'

  let devices: IDevice[] = [] // Initialize as an empty array

  const loadDevices = async () => {
    try {
      const { data } = await deviceService.getDevices()
      if (data && data.data) {
        devices = data.data
      }
    } catch (error) {
      console.error('Failed to load devices:', error)
    }
  }

  onMount(async () => await loadDevices())
</script>

<div class="grid grid-cols-1 gap-4 p-4 md:grid-cols-2 lg:grid-cols-3">
  {#each devices as device}
    <Card class="cursor-pointer rounded-lg  shadow-lg hover:bg-gray-100">
      <CardHeader>
        <CardTitle>{device.name}</CardTitle>
        <CardDescription>ID: {device.id}</CardDescription>
      </CardHeader>
      <CardContent>
        <p class="text-sm">Created At: {new Date(device.created_at).toLocaleString()}</p>
        <p class="text-sm">Updated At: {new Date(device.updated_at).toLocaleString()}</p>
      </CardContent>
      <CardFooter class="flex justify-between">
        <button class="text-blue-500 hover:underline">Edit</button>
        <button class="text-red-500 hover:underline">Delete</button>
      </CardFooter>
    </Card>
  {/each}
</div>
