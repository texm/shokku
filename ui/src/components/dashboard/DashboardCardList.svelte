<script>
  import { goto } from "$app/navigation";

  import Icon from "$common/Icon.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import DashboardCard from "./DashboardCard.svelte";

  export let contentType = "";
  export let states = {};
  export let query;

  let contentTypeTitle;
  $: contentTypeTitle =
    contentType.substring(0, 1).toUpperCase() + contentType.substring(1);
</script>

<QueryDataWrapper {query} action="loading {contentTypeTitle}s">
  <div class="flex flex-row gap-4 p-2 items-center">
    <span class="text-2xl font-bold text-base-content">{contentTypeTitle}s</span
    >
    <div class="flex-grow" />
    <div class="flex justify-end" class:hidden={$query.data.length === 0}>
      <button
        class="btn btn-secondary btn-sm gap-2"
        on:click={() => goto(`/${contentType}/new`)}
      >
        Create <Icon size="sm" type="plus" />
      </button>
    </div>
  </div>

  {#if $query.data.length > 0}
    <div class="flex flex-col gap-4">
      {#each $query.data as info}
        <DashboardCard {contentType} {info} />
      {/each}
    </div>
  {:else}
    <div class="card card-bordered w-auto shadow-lg hover:shadow-xl bg-neutral">
      <div class="card-body">
        <div class="grid grid-cols-2 items-center">
          <div class="h-full text-neutral-content">
            <span class="text-xl leading-8">No {contentTypeTitle}s</span>
          </div>
        </div>

        <div class="mt-4">
          <button
            class="btn gap-2 btn-primary"
            on:click={() => goto(`/${contentType}/new`)}
          >
            <Icon size="md" type="add" />
            <span class="text-lg">Create {contentTypeTitle}</span>
          </button>
        </div>
      </div>
    </div>
  {/if}
</QueryDataWrapper>
