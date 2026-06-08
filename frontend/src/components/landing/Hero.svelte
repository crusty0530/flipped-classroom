<script>
	import Header from "./Header.svelte";
	import SectionWrapper from "./SectionWrapper.svelte";

</script>
<SectionWrapper>
    <Header />
    <div class="flex flex-col gap-10 flex-1 items-center justify-center pb-10 md:pb-14">
        <h2 class="text-5xl sm:text-6xl md:text-7xl lg:text-8xl max-w-300 mx-auto w-full text-center">Teach <span class="text-slate-600 line-through">Traditionally</span><br/>
            with <span class="text-blue-400">Invertli</span>
        </h2>

        <p class="text-xl sm:text-2xl md:text-3xl text-center max-w-250 mx-auto w-full">
            <span class="text-blue-400">Invert</span> how your class learns by a research-backed flipped classroom platform
        </p>
        <button class="border border-blue-400 rounded-md p-2 text-2xl bg-blue-400 text-white hover:bg-blue-700 transition-all ease-in-out hover:cursor-pointer">Get Started</button>
    </div>
</SectionWrapper>