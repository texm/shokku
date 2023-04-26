<script>
	import {base} from "$app/paths";

	const carouselImages = [
		{alt: "dashboard screenshot", src: `${base}/images/dashboard-light.webp`},
	  {alt: "dashboard screenshot", src: `${base}/images/dashboard-dark.webp`},
		{alt: "app overview", src: `${base}/images/app-overview.webp`},
		{alt: "service overview", src: `${base}/images/service-overview.webp`},
	]

	const nextIdx = i => ((i + 1) % carouselImages.length);
	const prevIdx = i => (carouselImages - i)
</script>

<div class="mockup-window rounded-xl bg-base-200 w-auto h-fit shadow-lg shadow-secondary">
	<div class="w-full carousel rounded-box">
	{#each carouselImages as image, i}
		<div id="img{1 + i}" class="carousel-item relative w-full">
			<img src={image.src} class="w-full" alt={image.alt} />
			<div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
				<a href="#img{1 + prevIdx(i)}" class="btn btn-circle">❮</a>
				<a href="#img{1 + nextIdx(i)}" class="btn btn-circle">❯</a>
			</div>
		</div>
	{/each}
	</div>
</div>
