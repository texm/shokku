<script>
  import { createEventDispatcher } from "svelte";

  export let options;
  if (!options) options = { image: "" };

  const dispatch = createEventDispatcher();
  const checkOptionsValid = () => {
    dispatch("validityChange", { valid: !!options["image"] });
  };
</script>

<div class="flex flex-col gap-2">
  <label class="input-group input-group-md">
    <span class="w-auto">Docker Image</span>
    <input
      type="text"
      class="input input-md input-bordered flex-grow"
      on:change={checkOptionsValid}
      bind:value={options["image"]}
    />
  </label>
</div>
