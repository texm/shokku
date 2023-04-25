<script>
  import { useMutation } from "@sveltestack/svelte-query";
  import { setDockerRegistry } from "$lib/api";
  import Error from "$common/Error.svelte";

  let editing = false;

  export let server;
  let username;
  let password;

  const onSuccess = () => {
    editing = false;
  };
  const registryAuthMutation = useMutation(
    () => setDockerRegistry({ server, username, password }),
    { onSuccess }
  );
</script>

<div class:hidden={editing}>
  <button class="btn" on:click={() => (editing = true)}>
    Configure Docker Registry
  </button>
</div>

<div class:hidden={!editing} class="bg-base-100 p-2 w-fit rounded-lg">
  <div class="form-control w-full max-w-xs mb-4">
    <label class="label" for="server-input">
      <span class="label-text">Server</span>
    </label>
    <input
      id="server-input"
      type="text"
      placeholder="docker.io"
      bind:value={server}
      class="input input-bordered w-full max-w-xs"
    />

    <label class="label" for="username">
      <span class="label-text">Username</span>
    </label>
    <input
      id="username"
      type="text"
      placeholder="username"
      bind:value={username}
      class:input-warning={$registryAuthMutation.isError}
      class="input input-bordered w-full max-w-xs"
    />

    <label class="label" for="password">
      <span class="label-text">Password</span>
    </label>
    <input
      id="password"
      type="password"
      placeholder=""
      bind:value={password}
      class:input-warning={$registryAuthMutation.isError}
      class="input input-bordered w-full max-w-xs"
    />
  </div>

  <button
    class="btn btn-primary"
    on:click={$registryAuthMutation.mutate}
    class:loading={$registryAuthMutation.isLoading}
  >
    Save
  </button>
  <button
    class="btn btn-ghost"
    on:click={() => (editing = false)}
    disabled={$registryAuthMutation.isLoading}
  >
    Cancel
  </button>

  {#if $registryAuthMutation.isError}
    <Error
      action="authenticating with registry"
      error={$registryAuthMutation.error}
    />
  {/if}
</div>
