<script>
  import { createEventDispatcher } from "svelte";
  import KVEditor from "$common/KVEditor.svelte";

  export let key;
  export let value = [];
  export let enabled = true;

  export let labelText = "";
  export let inputText = "";

  const dispatch = createEventDispatcher();
  const onChange = ({ detail }) => {
    dispatch("changed", { key, enabled, value: detail });
  };
</script>

<div class="bg-neutral rounded-lg w-full p-2">
  <label class="label">
    <span class="label-text text-neutral-content">{labelText}</span>
  </label>
  <span class="w-fit">{inputText}</span>
  <KVEditor neutralButtons={true} vars={value} on:changed={onChange} />
</div>
