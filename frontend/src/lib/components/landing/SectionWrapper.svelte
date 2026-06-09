<script>
    let { children, id = undefined } = $props()
</script>

<section {id} class="min-h-screen flex flex-col px-4">
    <div class="flex flex-col flex-1 max-w-350 mx-auto w-full">
        {@render children()}
    </div>

</section>