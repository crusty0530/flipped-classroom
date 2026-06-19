<script>
	import SubmitButton from "$lib/components/landing/SubmitButton.svelte";

</script>
<label for="username">Username or Email:</label>
<input type="text" class="border-blue-400 border-2 hover:bg-blue-100 rounded-xl transition-all">

<label for="password">Password:</label>
<input type="password" class="border-blue-400 border-2 hover:bg-blue-100 rounded-xl transition-all">

<SubmitButton />

<p class="flex justify-between items-center w-full">Don't have an account? <a href="/auth/register" class="text-blue-400 px-2 py-1">Sign up</a></p>