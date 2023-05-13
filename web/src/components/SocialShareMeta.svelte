<script>
	export let title = "";
</script>

<svelte:head>
	<meta property="og:title" content={title}>
	<meta property="og:type" content="website" />
	<meta property="og:image" content="https://shokku.dev/share-card.png">
	<meta name="twitter:card" content="https://shokku.dev/share-card.png">

	<meta property="og:description" content="Free and open source interface for the Dokku PaaS">
	<meta name="twitter:image:alt" content="Shokku Logo">
</svelte:head>