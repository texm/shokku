<script>
  import { createEventDispatcher, onMount } from "svelte";

  export let options;
  if (!options) {
    options = {
      deployment_branch: "main",
      // envVar: "GIT_REV",
    };
  }

  const dispatch = createEventDispatcher();
  const checkOptionsValid = () => {
    dispatch("validityChange", { valid: options.deployment_branch !== "" });
  };
  onMount(() => dispatch("validityChange", { valid: true }));
</script>

<div class="flex flex-col gap-2">
  <label class="input-group input-group-md">
    <span class="w-auto">Deployment Branch</span>
    <input
      type="text"
      class="input input-md input-bordered flex-grow"
      bind:value={options["deployment_branch"]}
      on:change={checkOptionsValid}
    />
  </label>

  <!--label class="input-group input-group-md">
        <span class="w-auto">Git Revision Environment Variable</span>
        <input type="text" class="input input-md input-bordered flex-grow"
               bind:value={options["envVar"]} />
    </label-->
</div>
