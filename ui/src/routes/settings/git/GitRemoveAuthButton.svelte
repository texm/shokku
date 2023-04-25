<script>
  import { useMutation } from "@sveltestack/svelte-query";
  import { removeGitAuth } from "$lib/api";
  import Error from "$common/Error.svelte";

  let editing = false;

  let host;

  const onSuccess = () => (editing = false);
  const authMutation = useMutation(() => removeGitAuth({ host }), {
    onSuccess,
  });
</script>

<div class:hidden={editing}>
  <button class="btn" on:click={() => (editing = true)}>
    Remove Git Remote Server Credentials
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
  </div>

  <button
    class="btn btn-primary"
    on:click={$authMutation.mutate}
    class:loading={$authMutation.isLoading}
  >
    Remove
  </button>
  <button
    class="btn btn-ghost"
    on:click={() => (editing = false)}
    disabled={$authMutation.isLoading}
  >
    Cancel
  </button>

  {#if $authMutation.isError}
    <Error action="removing host" error={$authMutation.error} />
  {/if}
</div>
