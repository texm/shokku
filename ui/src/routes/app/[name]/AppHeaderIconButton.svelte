<script>
  import Icon from "$common/Icon.svelte";
  import { createEventDispatcher } from "svelte";

  export let action;
  export let loading;

  const dispatch = createEventDispatcher();
</script>

<button
  class="btn btn-outline md:btn-sm"
  class:gap-2={!loading}
  class:loading
  on:click={() => dispatch("clicked")}
>
  {#if !loading}
    <Icon type={action} size="sm" />
  {/if}
  {action}
</button>
