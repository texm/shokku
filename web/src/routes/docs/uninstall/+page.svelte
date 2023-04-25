<script>
	import Hint from "$components/Hint.svelte";

  const commands = [
		"dokku apps:destroy --force shokku",
		"dokku ssh-keys:remove shokkuadmin",
		"rm -r /var/lib/dokku/data/storage/shokku",
	]
</script>

<p>
	To completely uninstall Shokku, run the following commands
</p>

<div class="mockup-code mb-4">
	{#each commands as cmd, i}
		<pre data-prefix="{i}" class="m-0 my-1 p-0"><code>{cmd}</code></pre>
	{/each}
</div>

<Hint>
	This will not uninstall Dokku, or remove any apps/services
	you have created/deployed.
</Hint>

<h2>Delete Github App</h2>
<p>
	If you have completed setup with a github app, head to your
	<a href="https://github.com/settings/apps">Github Developer Settings</a>, click
	<i>Edit</i>, then <i>Advanced</i> in the sidebar, then <i>Delete Github App</i>.
</p>