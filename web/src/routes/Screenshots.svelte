<script>
	const carouselImages = [
		{alt: "dashboard screenshot", src: "/images/dashboard-light.webp"},
		{alt: "app overview", src: "/images/app-overview.webp"},
		{alt: "service overview", src: "/images/service-overview.webp"},
	]

	const nextIdx = i => ((i + 1) % carouselImages.length);
	const prevIdx = i => (carouselImages - i)
</script>

<div class="mockup-window rounded-xl bg-base-200 shadow-secondary w-3/5 h-auto">
	<img class="inline dark:hidden" alt="dashboard screenshot" src="/images/dashboard-dark.webp" />
	<img class="hidden dark:inline" alt="dashboard screenshot" src="/images/dashboard-light.webp" />
</div>
