<script lang="ts">
    const {name, pfp, href, subtitle} = $props();
</script>
<style>
    img {
        width: 20vw;
        object-fit: cover;
        aspect-ratio: 1;
        border-radius: 15px;
    }
</style>
<div>
    <a {href} target="_blank">
        <img src={pfp} alt={name + "'s headshot"} />
    </a>
    <h3 class="font-bold text-2xl mt-5 mb-1">{name}</h3>
    <div style="font-style: italic;">{subtitle}</div>
</div>