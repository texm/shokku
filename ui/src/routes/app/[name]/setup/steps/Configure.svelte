<script>
  import { createEventDispatcher } from "svelte";

  export let data = {};
  export let props = {};

  let sourceComponent;
  $: if (data.selectedSource && props.sources[data.selectedSource]) {
    sourceComponent = props.sources[data.selectedSource]?.component;
  }

  $: if (!data.sourceOptions) {
    data.sourceOptions = {};
  }

  const dispatch = createEventDispatcher();
  const onValidityChanged = ({ valid }) => {
    dispatch("statusChange", { complete: valid });
  };
</script>

<span class="text-lg">Configure source:</span>

{#if sourceComponent}
  <svelte:component
    this={sourceComponent}
    bind:options={data.sourceOptions[data.selectedSource]}
    on:validityChange={(e) => onValidityChanged(e.detail)}
  />
{/if}
