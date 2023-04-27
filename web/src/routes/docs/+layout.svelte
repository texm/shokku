<script>
	import {base} from "$app/paths";
	import {page} from "$app/stores";
	import Navbar from "$components/Navbar.svelte";
	import ProgressButtons from "./ProgressButtons.svelte";

	const sidebarPages = [
		{section: "Getting Started", ref: "/docs/intro", label: "Introduction"},
		{ref: "/docs/installation", label: "Install"},
		{ref: "/docs/setup", label: "Setup"},
		{section: "Usage", ref: "/docs/apps", label: "Apps"},
		{ref: "/docs/services", label: "Services"},
	  // {ref: "/docs/settings", label: "Settings"},
		{section: "Maintenance", ref: "/docs/upgrading", label: "Upgrading"},
		{ref: "/docs/uninstall", label: "Uninstall"},
	  {ref: "/docs/support", label: "Support"},
	];

	let currentPage, prevPage, nextPage;
	$: if ($page.route.id) {
		currentPage = sidebarPages.findIndex(p => p.ref === $page.route.id);
		prevPage = (currentPage > 0) ? sidebarPages[currentPage - 1] : null;
		nextPage = (currentPage < sidebarPages.length - 1) ? sidebarPages[currentPage + 1] : null;
	}
</script>

<svelte:head>
	<title>Shokku Docs - {sidebarPages[currentPage].label}</title>
</svelte:head>

<div class="flex flex-col h-screen">
	<div class="flex-shrink">
		<Navbar hideGetStarted={true}>
			<div slot="sidebar-toggle" class="flex-none lg:hidden">
				<label for="sidebar" class="btn btn-square btn-ghost drawer-button">
					<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
							 class="inline-block w-5 h-5 stroke-current">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
									d="M4 6h16M4 12h16M4 18h16"></path>
					</svg>
				</label>
			</div>
		</Navbar>
	</div>

	<div class="drawer drawer-mobile h-auto flex-grow">
		<input id="sidebar" type="checkbox" class="drawer-toggle" />

		<div class="drawer-content flex flex-col">
			<div class="p-8">
				<article class="prose">
					<h1 class="mb-0">{sidebarPages[currentPage].label}</h1>
					<hr class="mt-4 mb-8" />
					<slot />
				</article>
			</div>

			<div class="divider mt-2 mb-0"></div>

			<ProgressButtons previous={prevPage} next={nextPage} />
		</div>

		<div class="drawer-side">
			<label for="sidebar" class="drawer-overlay"></label>

			<ul class="menu w-80 bg-base-200 text-base-content pt-2">
				{#each sidebarPages as p, i}
					{#if p.section}
						<li class="menu-title py-1">
							<span>{p.section}</span>
						</li>
					{/if}
					<li class:bordered={currentPage === i}>
						<a href={base + p.ref}>
							<span class:pl-5={p.level === 1}>
								{p.label}
							</span>
						</a>
					</li>
				{/each}
			</ul>
		</div>
	</div>
</div>