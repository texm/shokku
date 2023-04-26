<script>
	import {base} from "$app/paths";

	const carouselImages = [
		{alt: "dashboard screenshot", src: `${base}/images/dashboard-light.webp`},
	  {alt: "dashboard screenshot", src: `${base}/images/dashboard-dark.webp`},
		{alt: "app overview", src: `${base}/images/app-overview.webp`},
		{alt: "service overview", src: `${base}/images/service-overview.webp`},
	]

	const len = carouselImages.length;
	const modIdx = i => (1 + (i + len) % len);
</script>

<div class="mockup-window rounded-xl bg-base-200 w-auto h-fit shadow-lg shadow-secondary">
	<div class="w-full carousel rounded-box">
	{#each carouselImages as image, i}
		<div id="img{modIdx(i)}" class="carousel-item relative w-full">
			<img src={image.src} class="w-full" alt={image.alt} />
			<div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
				<a href="#img{modIdx(i - 1)}" class="btn btn-circle">❮</a>
				<a href="#img{modIdx(i + 1)}" class="btn btn-circle">❯</a>
			</div>
		</div>
	{/each}
	</div>
</div>
