<script lang="ts">
    import { Button } from "@/components/ui/button";
  import {Input} from "@/components/ui/input"
</script>

<div class="flex place-items-center justify-center h-full">
    <div class="flex flex-col justify-center place-items-center">
        <Input placeholder="Username" class="m-3" />
        <Input placeholder="Password" class="m-3"  />
        <Button class="m-3">Login</Button>
    </div>
</div>