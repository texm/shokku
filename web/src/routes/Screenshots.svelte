<script>
	import {base} from "$app/paths";

	const carouselImages = [
		{alt: "dashboard screenshot", src: `${base}/images/dashboard-light.webp`},
		{alt: "app overview", src: `${base}/images/app-overview.webp`},
		{alt: "service overview", src: `${base}/images/service-overview.webp`},
	]

	const nextIdx = i => ((i + 1) % carouselImages.length);
	const prevIdx = i => (carouselImages - i)
</script>

<div class="mockup-window rounded-xl bg-base-200 w-auto h-fit shadow-lg shadow-secondary">
	<img class="inline dark:hidden" alt="dashboard screenshot" src="{base}/images/dashboard-dark.webp" />
	<img class="hidden dark:inline" alt="dashboard screenshot" src="{base}/images/dashboard-light.webp" />
</div>
