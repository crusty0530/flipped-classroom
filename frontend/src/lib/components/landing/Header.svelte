<script>
	import { goto } from "$app/navigation";

</script>
<header class="flex flex-col relative z-20">
    <div class="max-w-350 mx-auto w-full flex items-center justify-between p-4 py-6">
        <a href="/">
            <h1 class="font-semibold text-3xl group cursor-pointer">
                <span class="text-blue-400 group-hover:text-black transition-all">Invert</span><span class="text-black group-hover:text-blue-400 transition-all">li</span>
            </h1>    
        </a>
        <nav class="flex items-center md:gap-4 lg:gap-6">
            <a href="#WhatIs" class="duration-200 hover:text-blue-400 cursor-pointer">What is Invertli?</a>
            <a href="#WhoAre" class="duration-200 hover:text-blue-400 cursor-pointer">Who are We?</a>
            <button onclick={() => goto('/auth/register')} class="border border-blue-400 rounded-md p-2 hover:rounded-3xl hover:bg-blue-100 transition-all duration-300 ease-in-out">Get Started</button>
        </nav>
    </div>

</header>