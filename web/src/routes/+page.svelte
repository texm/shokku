<script>
	import SocialShareMeta from "$components/SocialShareMeta.svelte";
	import Navbar from "$components/Navbar.svelte";
	import Hero from "./Hero.svelte";
	import Features from "./Features.svelte";
	import Screenshots from "./Screenshots.svelte";
	import Footer from "./Footer.svelte";
</script>

<svelte:head>
	<title>Shokku</title>
	<SocialShareMeta title="Shokku" />
</svelte:head>

<div class="flex flex-col bg-base-300">
	<div class="flex-shrink">
		<Navbar />
	</div>

	<div class="flex-grow bg-base-100">
		<Hero />
	</div>

	<div class="flex flex-col md:flex-row justify-center p-8 gap-8">
		<div class="w-full md:w-3/5">
			<Screenshots />
		</div>

		<Features />
	</div>

	<div class="bg-base-200">
		<Footer />
	</div>
</div>