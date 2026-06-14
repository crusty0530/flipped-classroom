<script lang="ts">
    let { children } = $props();
</script>

<div class="grid h-screen place-items-center">
    <form action="submit" class="flex flex-col gap-4 border-5 border-blue-400 py-12 px-16 rounded-4xl">
        {@render children()}  
    </form>
</div>
