<script>
	import Navbar from "$components/Navbar.svelte";
  import Hero from "./Hero.svelte";
  import Features from "./Features.svelte";
</script>

<div class="flex flex-col h-screen">
	<div class="flex-shrink">
		<Navbar />
	</div>

	<div class="flex-grow bg-base-100">
		<Hero />
	</div>

	<div class="bg-base-300 justify-center p-8">
		<Features />
	</div>
</div>