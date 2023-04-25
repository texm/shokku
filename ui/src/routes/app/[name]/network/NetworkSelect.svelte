<script>
  import { createEventDispatcher } from "svelte";

  export let labelText;
  export let enable;
  export let selected = "";
  export let networks = [];
  export let loading;
  export let dirty;

  if (!selected) selected = "host";

  const dispatch = createEventDispatcher();
</script>

<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">{labelText}</span>
    <input type="checkbox" class="toggle" bind:checked={enable} />
  </label>
  {#if enable}
    <select
      bind:value={selected}
      class="select select-bordered w-full max-w-xs"
    >
      {#each networks as network}
        <option value={network}>{network}</option>
      {/each}
    </select>
  {/if}

  <button
    class="btn btn-primary my-2"
    on:click={() => dispatch("save")}
    class:hidden={!dirty}
    class:loading
  >
    Update
  </button>
</div>
