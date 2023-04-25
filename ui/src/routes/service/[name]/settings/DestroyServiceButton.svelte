<script>
  import { page } from "$app/stores";
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query";
  import { destroyService } from "$lib/api";
  import { goto } from "$app/navigation";
  import Error from "$common/Error.svelte";
  import ConfirmationModal from "$common/ConfirmationModal.svelte";
  import Icon from "../../../../components/common/Icon.svelte";

  export let serviceType;

  const serviceName = $page.params.name;
  let modalOpen = false;

  const queryClient = useQueryClient();
  const onSuccess = async () => {
    await queryClient.invalidateQueries("listServices");
    await goto("/");
  };
  const destroyMutation = useMutation(
    () => destroyService(serviceType, serviceName),
    { onSuccess }
  );
</script>

<div class="">
  <button
    class="btn btn-error gap-2 w-56"
    class:loading={$destroyMutation.isLoading}
    on:click={() => (modalOpen = true)}
  >
    {#if $destroyMutation.isLoading}
      Destroying...
    {:else}
      Destroy Service
      <Icon type="delete" />
    {/if}
  </button>
</div>

{#if $destroyMutation.isError}
  <div class="my-2">
    <Error action="destroying service" error={$destroyMutation.error} />
  </div>
{/if}

<ConfirmationModal
  name="destroy-service-modal"
  title="Destroy Service"
  action="destroy {serviceType} service {serviceName}"
  on:accepted={$destroyMutation.mutate}
  bind:open={modalOpen}
/>
