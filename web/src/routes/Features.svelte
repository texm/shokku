<script>
	const features = [
	  {title: "Free and Open Source", ref: "", description: "description"},
	  {title: "Github Integration", ref: "", description: "description"},
	  {title: "Services Support", ref: "", description: "description"},
	]
</script>

<div class="flex flex-col gap-2">
	{#each features as feature}
		<div class="card bg-base-100 shadow-xl">
			<div class="card-body">
				<h2 class="card-title">{feature.title}</h2>
				<!--p>{feature.description}</p-->
				<!--div class="card-actions justify-end">
					<button class="btn btn-ghost">Learn More</button>
				</div-->
			</div>
		</div>
	{/each}
</div>
