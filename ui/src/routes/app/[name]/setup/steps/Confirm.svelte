<script>
  import { createEventDispatcher, onMount } from "svelte";

  export let props = {};
  export let data = {};

  let selectedSource = data.selectedSource;
  let sourceOptions = data.sourceOptions[selectedSource];
  const optionDisplayText = (option) => {
    const value = sourceOptions[option];
    if (typeof value === "object") return value.name;
    return value;
  };

  const dispatch = createEventDispatcher();
  onMount(() => dispatch("statusChange", { complete: true }));
</script>

<span
  >Going to
  <span class="font-bold">{props.sources[selectedSource].createText}</span>
  with configuration</span
>
<div class="overflow-x-auto mt-3">
  <table class="table table-compact w-full">
    <thead>
      <tr><th>Key</th><th>Value</th></tr>
    </thead>
    <tbody>
      {#each Object.keys(sourceOptions) as option}
        {#if !option.startsWith("_") && optionDisplayText(option) !== ""}
          <tr>
            <td>{option}</td>
            <td>{optionDisplayText(option)}</td>
          </tr>
        {/if}
      {/each}
    </tbody>
  </table>
</div>
