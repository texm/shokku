<script>
  import Icon from "$common/Icon.svelte";

  export let cpuLimit;
  export let memLimit;
  export let memReserved;

  let showCPU, showMemLimit, showMemReserved;
  $: {
    showCPU = cpuLimit && cpuLimit["amount"];
    showMemLimit = memLimit && memLimit["amount"];
    showMemReserved = memReserved && memReserved["amount"];
  }
</script>

{#if showCPU || showMemLimit || showMemReserved}
  <div class="flex flex-row gap-4">
    {#if showCPU}
      <div class="stats shadow">
        <div class="stat">
          <div class="stat-title">vCPU Limit</div>
          <div class="stat-value items-center flex gap-2">
            <Icon type="cube" size="lg" />
            <span>{cpuLimit["amount"]}</span>
          </div>
        </div>
      </div>
    {/if}

    {#if showMemLimit}
      <div class="stats shadow">
        <div class="stat">
          <div class="stat-title">Memory Limit</div>
          <div class="stat-value items-center flex gap-2">
            <Icon type="layers" size="lg" />
            <span>{memLimit["amount"]}{memLimit["unit"]}</span>
          </div>
        </div>
      </div>
    {/if}

    {#if showMemReserved}
      <div class="stats shadow">
        <div class="stat">
          <div class="stat-title">Memory Reserved</div>
          <div class="stat-value items-center flex gap-2">
            <Icon type="layers" size="lg" />
            <span>{memReserved["amount"]}{memReserved["unit"]}</span>
          </div>
        </div>
      </div>
    {/if}
  </div>
{/if}
