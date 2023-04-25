<script>
  import { createEventDispatcher } from "svelte";

  export let key;
  export let showToggle = true;
  export let enabled = false;
  export let type = "string";
  export let value = "";

  export let toggleText = "";
  export let labelText = "";
  export let inputText = "";

  const dispatch = createEventDispatcher();
  const onChange = () =>
    dispatch("changed", { key, enabled, value: value.toString() });
</script>

<div class="bg-neutral rounded-lg w-full p-2">
  {#if showToggle}
    <label class="label cursor-pointer">
      <span class="label-text text-neutral-content">{toggleText}</span>
      <input
        type="checkbox"
        class="toggle"
        bind:checked={enabled}
        on:change={onChange}
      />
    </label>
  {/if}
  <div class:hidden={showToggle && !enabled}>
    <label class="label" class:hidden={!labelText}>
      <span class="label-text text-neutral-content">{labelText}</span>
    </label>
    <label class="input-group w-full flex">
      <span class="w-fit">{inputText}</span>
      {#if type === "number"}
        <input
          bind:value
          type="number"
          class="input input-bordered flex-grow"
          on:change={onChange}
        />
      {:else}
        <input
          bind:value
          class="input input-bordered flex-grow"
          on:change={onChange}
        />
      {/if}
    </label>
  </div>
</div>
