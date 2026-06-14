<script>
	import Hero from "$lib/components/landing/Hero.svelte";
	import WhatIs from "$lib/components/landing/WhatIs.svelte";
	import WhoAre from "$lib/components/landing/WhoAre.svelte";
    import Footer from "$lib/components/landing/Footer.svelte";

</script>
<div class="inset-0 opacity-5 absolute z-[-1] bgGrid"></div>

<main class="flex flex-col">
    <Hero />
    <WhatIs />
    <WhoAre />
    <Footer />
</main>

<style>
    .bgGrid {
    background-size: 40px 40px;
    background-image: linear-gradient(to right, grey 1px, transparent 1px), linear-gradient(to bottom, grey 1px, transparent 1px)
}
</style>