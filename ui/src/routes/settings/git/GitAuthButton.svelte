<script>
  import { useMutation } from "@sveltestack/svelte-query";
  import { addGitAuth } from "$lib/api";
  import Error from "$common/Error.svelte";

  let editing = false;

  let host;
  let username;
  let password;

  const onSuccess = () => (editing = false);
  const authMutation = useMutation(
    () => addGitAuth({ host, username, password }),
    { onSuccess }
  );
</script>

<div class:hidden={editing}>
  <button class="btn" on:click={() => (editing = true)}>
    Add Git Remote Server Credentials
  </button>
</div>

<div class:hidden={!editing} class="bg-base-100 p-2 w-fit rounded-lg">
  <div class="form-control w-full max-w-xs mb-4">
    <label class="label" for="server-input">
      <span class="label-text">Host</span>
    </label>
    <input
      id="server-input"
      type="text"
      placeholder="github.com"
      bind:value={host}
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
      class:input-warning={$authMutation.isError}
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
      class:input-warning={$authMutation.isError}
      class="input input-bordered w-full max-w-xs"
    />
  </div>

  <button
    class="btn btn-primary"
    on:click={$authMutation.mutate}
    class:loading={$authMutation.isLoading}
  >
    Save
  </button>
  <button
    class="btn btn-ghost"
    on:click={() => (editing = false)}
    disabled={$authMutation.isLoading}
  >
    Cancel
  </button>

  {#if $authMutation.isError}
    <Error action="authenticating with host" error={$authMutation.error} />
  {/if}
</div>
