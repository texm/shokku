<script>
  import { getGlobalDomainsList, addGlobalDomain } from "$lib/api";

  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Icon from "$common/Icon.svelte";
  import Modal from "$common/Modal.svelte";

  import DomainListItem from "./DomainListItem.svelte";
  import {
    useMutation,
    useQuery,
    useQueryClient,
  } from "@sveltestack/svelte-query";
  import Card from "$common/Card.svelte";

  const queryClient = useQueryClient();
  const globalDomains = useQuery("getGlobalDomainsList", getGlobalDomainsList);
  const invalidateDomains = () =>
    queryClient.invalidateQueries("getGlobalDomainsList");

  let addDomainModalOpen = false;

  let newDomain;
  const addDomainMutation = useMutation(() => addGlobalDomain(newDomain), {
    onSuccess: () => {
      invalidateDomains();
      addDomainModalOpen = false;
    },
  });
</script>

<QueryDataWrapper query={globalDomains} action="loading domains">
  <Card title="Global Domains">
    <div class="flex flex-col gap-3 w-72">
      {#each $globalDomains.data as domain, i}
        <div class="">
          <DomainListItem {domain} on:removed={invalidateDomains} />
        </div>
      {/each}
      {#if $globalDomains.data.length === 0}
        <span>No domains registered</span>
      {/if}
    </div>

    <div slot="actions">
      <button class="btn gap-3" on:click={() => (addDomainModalOpen = true)}>
        New Domain
        <Icon type="add" />
      </button>
    </div>
  </Card>
</QueryDataWrapper>

<Modal
  name="add-domain"
  title="Add a new global domain"
  bind:open={addDomainModalOpen}
  preventClose={$addDomainMutation.isLoading}
>
  <div class="form-control">
    <label class="input-group input-group-vertical">
      <span>Domain</span>
      <input
        type="text"
        placeholder="example.com"
        class="input input-bordered"
        bind:value={newDomain}
        disabled={$addDomainMutation.isLoading}
      />
    </label>
  </div>

  <div class="mt-4">
    <button
      class="btn"
      class:loading={$addDomainMutation.isLoading}
      on:click={$addDomainMutation.mutate}
    >
      Confirm
    </button>
  </div>
</Modal>
