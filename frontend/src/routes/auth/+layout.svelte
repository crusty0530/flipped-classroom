<script lang="ts">
    let { children } = $props();
</script>



<div class="grid h-screen place-items-center">
    <div class="flex flex-col gap-10 flex-1 items-center justify-center pb-10 md:pb-14">
        <h2 class="text-5xl sm:text-6xl md:text-7xl lg:text-8xl max-w-300 mx-auto w-full text-center font-bold">
            <span class="text-blue-400">Invert</span>li
        </h2>
    </div>
    
    <form action="submit" class="flex flex-col gap-4 py-12 px-16 rounded-4xl shadow-2xl shadow-black">
        {@render children()}  
    </form>
</div>
