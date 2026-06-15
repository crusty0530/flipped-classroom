<script>
	import { goto } from "$app/navigation";
	import Header from "./Header.svelte";
	import SectionWrapper from "./SectionWrapper.svelte";
</script>
<SectionWrapper>
    <Header />
    <div class="flex flex-col gap-10 flex-1 items-center justify-center pb-10 md:pb-14">
        <h2 class="text-5xl sm:text-6xl md:text-7xl lg:text-8xl max-w-300 mx-auto w-full text-center">Teach <span class="text-slate-600 line-through">Traditionally</span><br/>
            with <span class="text-blue-400 hover:text-black transition-all">Invertli</span>
        </h2>

        <p class="text-xl sm:text-2xl md:text-3xl text-center max-w-300 mx-auto w-full">
            <span class="text-blue-400 hover:text-black transition-all">Invert</span> how your class learns with a research-backed flipped classroom platform
        </p>
        <button onclick={() => goto('auth/register')} class="border border-blue-400 rounded-md p-2 text-2xl bg-blue-400 text-white hover:bg-blue-700 hover:rounded-3xl transition-all hover:cursor-pointer">Get Started</button>
    </div>
</SectionWrapper>