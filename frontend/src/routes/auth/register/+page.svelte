<script>
	import SubmitButton from "$lib/components/landing/SubmitButton.svelte";

</script>
<label for="username" class="">Username:</label>
<input type="text" class="border-blue-400 border-2 hover:bg-blue-100 rounded-xl transition-all">

<label for="email">Email:</label>
<input type="email" class="border-blue-400 border-2 hover:bg-blue-100 rounded-xl transition-all">

<label for="password">Password:</label>
<input type="password" class="border-blue-400 border-2 hover:bg-blue-100 rounded-xl transition-all">

<SubmitButton />