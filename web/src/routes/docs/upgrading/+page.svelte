<script>
	import Hint from "$components/Hint.svelte";

	// TODO: update this
	const shokkuImage = "repo/shokku";
  const manualCmds = [
	  `export SHOKKU_VERSION=latest`,
	  `docker pull "${shokkuImage}:$SHOKKU_VERSION"`,
	  `dokku git:from-image shokku $(docker inspect --format='{{index .RepoDigests 0}}' "${shokkuImage}:$SHOKKU_VERSION")`
  ];
</script>

<Hint>
	Due to how Dokku tracks app changes, tracking the <i>latest</i> image tag will result
	in no changes being detected. Instead we use the image digest.
</Hint>

<div class="mockup-code mt-4">
	{#each manualCmds as cmd, i}
		<pre data-prefix="{i}" class="m-0 my-1 p-0"><code>{cmd}</code></pre>
	{/each}
</div>