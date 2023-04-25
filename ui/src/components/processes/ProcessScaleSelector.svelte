<script>
  import { page } from "$app/stores";
  import { setAppProcessScale } from "$lib/api";
  import { commandExecutionIds, executionIdDescriptions } from "$lib/stores";
  import { createEventDispatcher } from "svelte";
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query";

  export let processName;
  export let scale;

  const appName = $page.params.name;
  const queryClient = useQueryClient();

  let scaleRange = [];
  $: if (scale !== null) {
    let scaleRangeStart = Math.max(0, scale - 3);
    scaleRange = [...Array(6).keys()].map((i) => i + scaleRangeStart);
  }

  const dispatch = createEventDispatcher();

  const invalidateAppOverview = () =>
    queryClient.invalidateQueries(["getAppOverview", appName]);

  const setScaleMutation = useMutation(
    async (newScale) => {
      const id = await setAppProcessScale(appName, processName, newScale);
      $executionIdDescriptions[
        id
      ] = `Scaling ${appName}-${processName} [${scale}->${newScale}]`;
      return commandExecutionIds.addID(id);
    },
    {
      onSuccess: (_, newScale) => {
        if (newScale === 0 || scale === 0) invalidateAppOverview();
        scale = newScale;
      },
    }
  );
</script>

<div class="flex flex-col gap-2">
  <div class="items-center flex flex-row gap-2">
    <span class="text-lg">Scale: </span>
    {#if $setScaleMutation.isLoading}
      <button class="btn btn-square btn-sm loading" />
    {:else}
      <button
        class="btn btn-square btn-sm"
        on:click={() => $setScaleMutation.mutate(scale - 1)}
      >
        -
      </button>
      <select
        class="select select-sm select-ghost"
        disabled={$setScaleMutation.isLoading}
        on:change={({ target }) =>
          $setScaleMutation.mutate(parseInt(target.value))}
      >
        {#each scaleRange as n}
          <option value={n} selected={scale === n}>{n}</option>
        {/each}
      </select>
      <button
        class="btn btn-square btn-sm"
        on:click={() => $setScaleMutation.mutate(scale + 1)}
      >
        +
      </button>
    {/if}
  </div>
</div>
