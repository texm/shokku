<script>
  import { goto } from "$app/navigation";
  import { destroyApp } from "$lib/api";
  import { page } from "$app/stores";

  import ConfirmationModal from "$common/ConfirmationModal.svelte";
  import Icon from "$common/Icon.svelte";
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query";
  import Error from "$common/Error.svelte";

  const appName = $page.params.name;

  let confirmationModalOpen = false;

  const queryClient = useQueryClient();
  const onSuccess = async () => {
    await queryClient.invalidateQueries("getAllAppsOverview");
    await goto("/");
  };
  const destroyAppMutation = useMutation(destroyApp, { onSuccess });
</script>

<div class="">
  <button
    class="btn btn-error gap-2 w-56"
    class:loading={$destroyAppMutation.isLoading}
    on:click={() => (confirmationModalOpen = true)}
  >
    {#if $destroyAppMutation.isLoading}
      Destroying...
    {:else}
      Destroy App
      <Icon type="delete" />
    {/if}
  </button>
</div>

{#if $destroyAppMutation.isError}
  <div class="my-2">
    <Error action="destroying app" error={$destroyAppMutation.error} />
  </div>
{/if}

<ConfirmationModal
  name="destroy-app-modal"
  title="Destroy App"
  action="destroy {appName}"
  on:accepted={() => $destroyAppMutation.mutate(appName)}
  bind:open={confirmationModalOpen}
/>
