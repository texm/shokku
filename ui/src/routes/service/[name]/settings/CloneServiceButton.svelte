<script>
  import { cloneService } from "$lib/api";
  import { useMutation } from "@sveltestack/svelte-query";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";

  import Modal from "$common/Modal.svelte";
  import Icon from "$common/Icon.svelte";
  import Error from "$common/Error.svelte";

  const serviceName = $page.params.name;

  let newName = serviceName;

  let modalOpen = false;

  const cloneMutation = useMutation(
    (newName) => cloneService(serviceName, newName),
    { onSuccess: (_, newName) => goto(`/service/${newName}`) }
  );
</script>

<div class="">
  <button class="btn btn-info gap-2 w-56" on:click={() => (modalOpen = true)}>
    Clone Service
    <Icon type="edit" />
  </button>
</div>

{#if $cloneMutation.isError}
  <div class="my-2">
    <Error action="cloning service" error={$cloneMutation.error} />
  </div>
{/if}

<Modal
  name="create-app"
  title="Clone '{serviceName}'"
  bind:open={modalOpen}
  preventClose={$cloneMutation.isLoading}
>
  <div class="form-control w-full max-w-xs mb-4">
    <label class="label" for="new-name">
      <span class="label-text">Enter a new name</span>
    </label>
    <input
      id="new-name"
      type="text"
      bind:value={newName}
      disabled={$cloneMutation.isLoading}
      class="input input-bordered w-full max-w-xs"
    />
  </div>

  <button
    class="btn btn-primary"
    on:click={() => $cloneMutation.mutate(newName)}
    class:loading={$cloneMutation.isLoading}
  >
    Save
  </button>
  <button
    class="btn btn-ghost"
    on:click={() => (modalOpen = false)}
    disabled={$cloneMutation.isLoading}
  >
    Cancel
  </button>
</Modal>
