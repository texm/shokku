<script>
	import { base } from "$app/paths";

	export let next;
	export let previous;
</script>

<div class="flex w-full justify-between pl-6 pr-10 my-6">
	<div>
		{#if previous}
			<a href={base + previous.ref} class="btn btn-md md:btn-lg btn-ghost gap-2 normal-case lg:gap-3">
				<svg class="h-8 w-8 fill-current md:h-8 md:w-8" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M15.41,16.58L10.83,12L15.41,7.41L14,6L8,12L14,18L15.41,16.58Z"></path></svg>
				<div class="flex flex-col items-start">
					<span class="text-base-content/50 hidden text-xs font-normal md:block">
						Prev
					</span>
					<span>{previous.label}</span>
				</div>
			</a>
		{/if}
	</div>
	<div>
		{#if next}
			<a href={base + next.ref} class="btn btn-md md:btn-lg gap-2 normal-case lg:gap-3">
				<div class="flex flex-col items-start">
					<span class="text-neutral-content/50 hidden text-xs font-normal md:block">
						Next
					</span>
					<span>{next.label}</span>
				</div>
				<svg class="h-8 w-8 fill-current md:h-8 md:w-8" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z"></path></svg>
			</a>
		{/if}
	</div>
</div>