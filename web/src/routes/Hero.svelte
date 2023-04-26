<script>
	import {base} from "$app/paths";
  import Icons from "./Icons.svelte";

	const demoUrl = "https://demo.example.com"
</script>

<div class="hero h-full">
	<div class="hero-content text-center flex-col py-10 px-0">
		<div class="max-w-md mb-4">
			<h1 class="text-5xl font-bold mb-2">A friendly interface for Dokku</h1>
			<p class="text-xl px-4">
				an open-source self-hostable web UI to help
				<strong>configure</strong> and <strong>deploy</strong>
				your apps and services with <a class="link" href="https://dokku.com/">Dokku</a>
			</p>

			<div class="flex flex-col gap-4 mb-4 p-4">
				<Icons />
			</div>

			<div class="flex flex-row gap-4 justify-center">
				<a href="{base}/docs/intro">
					<button class="btn btn-primary text-primary-content">Get Started</button>
				</a>

				<!--a href={demoUrl}>
					<button class="btn btn-secondary text-secondary-content">Try the Demo</button>
				</a-->
			</div>
		</div>
	</div>
</div>