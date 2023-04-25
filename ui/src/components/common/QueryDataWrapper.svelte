<script>
  import Error from "./Error.svelte";
  import Loader from "./Loader.svelte";

  export let query;
  export let action = "";
</script>

{#if $query.isLoading}
  <Loader />
{:else if $query.isError}
  <Error error={$query.error} {action} />
{:else}
  <slot />
{/if}
