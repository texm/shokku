<script>
	const features = [
	  {title: "100% Free and Open Source", description: "free as in beer"},
	  {title: "Github Integration", description: "secure authentication for single users or organisations"},
	  {title: "Easy Setup", description: "ready to go in minutes"},
	]
</script>

<div class="flex flex-col gap-4">
	{#each features as feature}
		<div class="card bg-base-100 shadow-xl max-w-md">
			<div class="card-body">
				<h2 class="card-title">{feature.title}</h2>
				<p>{feature.description}</p>
				<!--div class="card-actions justify-end">
					<button class="btn btn-ghost">Learn More</button>
				</div-->
			</div>
		</div>
	{/each}
</div>
