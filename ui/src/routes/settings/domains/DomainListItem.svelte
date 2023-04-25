<script>
  import { removeGlobalDomain } from "$lib/api";
  import { createEventDispatcher } from "svelte";

  import Icon from "$common/Icon.svelte";
  import ConfirmationModal from "$common/ConfirmationModal.svelte";
  import Alert from "$common/Alert.svelte";
  import { useMutation } from "@sveltestack/svelte-query";
  import Error from "../../../components/common/Error.svelte";

  export let domain;

  let confirmationModalOpen = false;

  const dispatch = createEventDispatcher();

  const removeDomainMutation = useMutation(() => removeGlobalDomain(domain), {
    onSuccess: () => {
      dispatch("removed");
    },
  });
</script>

<div class="flex flex-row items-center gap-3 w-auto bg-base-100 p-2 rounded-lg">
  <div>
    <span class="text-lg">{domain}</span>
  </div>
  <div class="flex-grow" />
  <div class="">
    <button
      class="btn bg-base-100 hover:btn-error btn-sm btn-outline gap-2 btn-circle"
      class:loading={$removeDomainMutation.isLoading}
      on:click={() => (confirmationModalOpen = true)}
    >
      {#if !$removeDomainMutation.isLoading}
        <Icon type="delete" size="sm" />
      {/if}
    </button>
  </div>
</div>

{#if $removeDomainMutation.isError}
  <Error action="removing domain" error={$removeDomainMutation.error} />
{/if}

<ConfirmationModal
  name="remove-domain-modal"
  title="Removing Domain"
  action="remove domain {domain}"
  on:accepted={$removeDomainMutation.mutate}
  bind:open={confirmationModalOpen}
  doingAction={$removeDomainMutation.isLoading}
/>
