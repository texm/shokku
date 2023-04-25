<script>
  import { createEventDispatcher } from "svelte";

  export let props = {};
  export let data = {};

  const dispatch = createEventDispatcher();

  const sourceSelected = (source) => {
    data.selectedSource = source;
    dispatch("statusChange", { complete: true });
  };
</script>

<span class="text-lg">Choose a source:</span>
<div class="tabs tabs-boxed bg-base-100 mt-4">
  {#each Object.keys(props.sources) as source}
    <a
      class="tab"
      class:tab-active={source === data.selectedSource}
      on:click={() => sourceSelected(source)}
    >
      {props.sources[source].label}
    </a>
  {/each}
</div>
