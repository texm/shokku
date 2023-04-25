<script>
  import { page } from "$app/stores";
  import { removeAppDomain } from "$lib/api";
  import { createEventDispatcher } from "svelte";

  import Icon from "$common/Icon.svelte";
  import ConfirmationModal from "$common/ConfirmationModal.svelte";
  import Alert from "$common/Alert.svelte";

  export let domain = "";
  export let loading = false;

  let domainHref = domain.startsWith("http") ? domain : `https://${domain}`;
  const dispatch = createEventDispatcher();
  const signalRemoved = () => dispatch("removeDomain", { domain });

  let confirmationModalOpen = false;
</script>

<div
  class="flex flex-row items-center gap-2 w-auto shadow-lg bg-base-100 p-2 rounded-lg"
>
  <div>
    <a class="link link-hover" href={domainHref}
      ><span class="text-lg">{domain}</span></a
    >
  </div>
  <div class="flex-grow" />
  <div class="">
    <button
      class="btn bg-base-100 hover:btn-error btn-sm btn-outline gap-2 btn-circle"
      class:loading
      on:click={() => (confirmationModalOpen = true)}
    >
      {#if !loading}
        <Icon type="delete" size="sm" />
      {/if}
    </button>
  </div>
</div>

<ConfirmationModal
  name="remove-domain-modal"
  title="Removing Domain"
  action="remove domain {domain}"
  on:accepted={signalRemoved}
  bind:open={confirmationModalOpen}
  doingAction={loading}
/>
