<script>
  export let config;

  config = {
    alias: "",
    query_string: "",
  };
</script>

<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">Alternative name for environment variable</span>
  </label>
  <input
    bind:value={config["alias"]}
    type="text"
    placeholder="BLUE_DATABASE"
    class="input input-bordered w-full max-w-xs"
  />
</div>

<div class="form-control w-full max-w-xs">
  <label class="label">
    <span class="label-text">Query string to append to connection</span>
  </label>
  <input
    bind:value={config["query_string"]}
    type="text"
    placeholder="pool=5&foo=bar"
    class="input input-bordered w-full max-w-xs"
  />
</div>
