<script lang="ts">
	const { children } = $props();
</script>

<div class="people">
	{@render children()}
</div>

<style>
	.people {
		display: flex;
		justify-content: center;
		align-items: center;
		flex-wrap: nowrap;
		text-align: center;
		gap: 10%;
	}
</style>
