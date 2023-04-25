<script>
  import { getPlugins } from "$lib/api";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import { useQuery } from "@sveltestack/svelte-query";
  import Plugin from "./Plugin.svelte";
  import Card from "$common/Card.svelte";

  const plugins = useQuery("getServerPlugins", getPlugins);
  let showCore = false;
</script>

<QueryDataWrapper query={plugins} action="loading plugins">
  <Card title="Plugins">
    <div class="form-control">
      <label class="label cursor-pointer w-52">
        <span class="label-text">Show core plugins?</span>
        <input type="checkbox" class="toggle" bind:checked={showCore} />
      </label>
    </div>

    <div class="flex flex-wrap gap-2">
      {#each $plugins.data as pluginInfo}
        {#if !pluginInfo.description.startsWith("dokku core") || showCore}
          <Plugin {...pluginInfo} />
        {/if}
      {/each}
    </div>
  </Card>
</QueryDataWrapper>
