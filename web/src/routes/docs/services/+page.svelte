<script>
	const configOptions = [
	  {name: "Config Options", default: "", description: "Extra CLI arguments to pass to the container creation"},
	  {name: "Environment Variables", default: "", description: "Environment variables to start the service with"},
	  {name: "Image", default: "", description: "Custom image name to use"},
	  {name: "Version", default: "latest", description: "Image version to use"},
	  {name: "Memory", default: "", description: "Container memory limit in MB"},
	  {name: "Shared Memory", default: "", description: "Override shared memory size for container in MB"},
	  {name: "Password", default: "", description: "Override user-level service password"},
	  {name: "Root Password", default: "", description: "Override root-level service password"},
	]
</script>

<h2>Creating a service</h2>
<p>
	Currently you can create 4 types of service:
	<strong>MySQL</strong>, <strong>PostgreSQL</strong>,
	<strong>MongoDB</strong>, and <strong>Redis</strong>.

	They all support the same configuration options.
</p>

<h3>Configuration</h3>
<div class="overflow-x-auto">
	<table class="table w-full">
		<thead>
			<tr>
				<th>Option</th>
				<th>Default</th>
				<th>Description</th>
			</tr>
		</thead>
		<tbody>
		{#each configOptions as opt}
			<tr>
				<td>{opt.name}</td>
				<td>{opt.default}</td>
				<td>{opt.description}</td>
			</tr>
		{/each}
		</tbody>
	</table>
</div>

<h2>Linking a service</h2>
<p>
	Services need to be <i>linked</i> to an app. Once this is done, the connection
	information will be available in the UI, and injected as an environment variable
	to the app (eg <code>REDIS_URL</code>)
</p>