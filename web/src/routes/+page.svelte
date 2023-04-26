<script>
	import Navbar from "$components/Navbar.svelte";
  import Hero from "./Hero.svelte";
  import Features from "./Features.svelte";
  import Screenshots from "./Screenshots.svelte";
</script>

<div class="flex flex-col h-screen">
	<div class="flex-shrink">
		<Navbar />
	</div>

	<div class="flex-grow bg-base-100">
		<Hero />
	</div>

	<div class="bg-base-300 flex flex-col md:flex-row justify-center p-8 gap-8">
		<div class="w-full md:w-3/5">
			<Screenshots />
		</div>

		<Features />
	</div>
</div>