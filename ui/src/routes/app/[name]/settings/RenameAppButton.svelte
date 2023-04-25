<script>
  import { renameApp } from "$lib/api";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";

  import Modal from "$common/Modal.svelte";
  import Icon from "$common/Icon.svelte";
  import { useMutation } from "@sveltestack/svelte-query";

  const appName = $page.params.name;

  let newName = appName;
  let renameModalOpen = false;

  const renameMutation = useMutation((newName) => renameApp(appName, newName), {
    onSuccess: () => goto(`/app/${newName}`),
  });
</script>

<div class="">
  <button
    class="btn btn-info gap-2 w-56"
    on:click={() => (renameModalOpen = true)}
  >
    Rename App
    <Icon type="edit" />
  </button>
</div>

<Modal
  name="create-app"
  title="Rename '{appName}'"
  bind:open={renameModalOpen}
  preventClose={$renameMutation.isLoading}
>
  <div class="form-control w-full max-w-xs mb-4">
    <label class="label" for="new-name">
      <span class="label-text">Enter a new name</span>
    </label>
    <input
      id="new-name"
      type="text"
      bind:value={newName}
      disabled={$renameMutation.isLoading}
      class="input input-bordered w-full max-w-xs"
    />
  </div>

  <button
    class="btn btn-primary"
    on:click={() => $renameMutation.mutate(newName)}
    class:loading={$renameMutation.isLoading}
  >
    Save
  </button>
  <button
    class="btn btn-ghost"
    on:click={() => (renameModalOpen = false)}
    disabled={$renameMutation.isLoading}
  >
    Cancel
  </button>
</Modal>
