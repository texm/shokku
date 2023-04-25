<script>
  import { createEventDispatcher } from "svelte";

  export let options;
  if (!options) {
    options = {
      repository_url: "",
      git_ref: "",
      _use_custom_dockerfile_path: false,
      custom_dockerfile_path: "",
    };
  }

  const dispatch = createEventDispatcher();
  const checkOptionsValid = () => {
    // TODO: check repo access etc
    dispatch("validityChange", { valid: !!options["repository_url"] });
  };
</script>

<div class="flex flex-col gap-2">
  <label class="input-group input-group-md">
    <span class="w-auto">Repository URL</span>
    <input
      type="text"
      class="input input-md input-bordered flex-grow"
      placeholder="https://github.com/heroku/node-js-getting-started.git"
      bind:value={options["repository_url"]}
      on:change={checkOptionsValid}
    />
  </label>

  <span class="text-sm"
    >If the repository is private, ensure you set up access using the
    credentials found in
    <a href="/settings" class="link" target="_blank">settings</a>.
  </span>

  <label class="input-group input-group-md">
    <span class="w-auto">Git Reference (branch, tag, commit)</span>
    <input
      type="text"
      class="input input-md input-bordered flex-grow"
      placeholder="optional"
      bind:value={options["git_ref"]}
      on:change={checkOptionsValid}
    />
  </label>

  <!--div class="form-control w-60 rounded-lg">
    <label class="label cursor-pointer">
      <span class="label-text">Specify Dockerfile Path?</span>
      <input
        type="checkbox"
        class="checkbox checkbox-accent"
        bind:checked={options["_use_custom_dockerfile_path"]}
        on:change={checkOptionsValid}
      />
    </label>
  </div>

  {#if options["_use_custom_dockerfile_path"]}
    <label class="input-group input-group-md">
      <span class="w-auto">Dockerfile Path</span>
      <input
        type="text"
        class="input input-md input-bordered flex-grow"
        placeholder="Dockerfile"
        on:change={checkOptionsValid}
        bind:value={options["custom_dockerfile_path"]}
      />
    </label>
  {/if}
  !-->
</div>
