<script>
	import SectionWrapper from "./SectionWrapper.svelte";
	import People from "./who/People.svelte";
	import Person from "./who/Person.svelte";

    import topher from "$lib/assets/topher.jpg";
    import mikey from "$lib/assets/mikey.jpg";
</script>
<SectionWrapper id="WhoAre">
    <div class="flex flex-col gap-10 flex-1 items-center justify-center pb-10 md:pb-14">
        <h2 class="text-5xl sm:text-6xl md:text-7xl lg:text-8xl max-w-300 mx-auto w-full text-center">Who are <span class="text-blue-400">We</span>?</h2>

        <p class="text-xl sm:text-2xl md:text-3xl text-center max-w-250 mx-auto w-full">
            We are two college kids that want to fill the gaps in the field of public education and make tools more accessible to teachers.
        </p>
        
        <People>
            <Person name="Mikey Evarts" pfp={mikey} href="https://mikeylab.com" subtitle="Mikey"></Person>
            <Person name="Christopher White" pfp={topher} href="https://www.linkedin.com/in/christopher-white-cs/" subtitle="Topher"></Person>
        </People>
    </div>
</SectionWrapper>