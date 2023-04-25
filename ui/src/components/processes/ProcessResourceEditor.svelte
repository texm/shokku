<script>
  import { setAppProcessResources } from "$lib/api";

  import { createEventDispatcher } from "svelte";
  import { page } from "$app/stores";
  import { useMutation } from "@sveltestack/svelte-query";
  import Error from "$common/Error.svelte";
  import ProcessResource from "./ProcessResource.svelte";

  export let processName;
  export let resourceDefaults;

  export let cpuLimit;
  export let memLimit;
  export let memReserved;

  const appName = $page.params.name;

  const setResourcesMutation = useMutation((resources) =>
    setAppProcessResources(appName, processName, resources)
  );
  const dispatch = createEventDispatcher();

  let defaultCPULimit, defaultMemLimit, defaultMemReservation;
  const setDefaults = () => {
    const limits = resourceDefaults["limits"] || {};
    const reservations = resourceDefaults["reservations"] || {};
    defaultCPULimit = limits["cpu"];
    defaultMemLimit = limits["memory"];
    defaultMemReservation = reservations["memory"];
  };
  $: if (resourceDefaults) setDefaults();

  let cpuLimitAmount;
  let memLimitAmount, memLimitUnit;
  let memReservedAmount, memReservedUnit;
  const setValues = () => {
    cpuLimitAmount = cpuLimit["amount"];
    memLimitAmount = memLimit["amount"];
    memLimitUnit = memLimit["unit"];
    memReservedAmount = memReserved["amount"];
    memReservedUnit = memReserved["unit"];
  };
  $: if (cpuLimit || memLimit || memReserved) setValues();

  const submit = async () => {
    const limits = {
      cpu: cpuLimitAmount,
      memory: memLimitAmount,
      memory_unit: memLimitUnit,
    };
    const reservations = {
      memory: memReservedAmount,
      memory_unit: memReservedUnit,
    };
    await $setResourcesMutation.mutateAsync({ limits, reservations });
    dispatch("successfulEdit");
  };
</script>

<div class="text-neutral">
  <div class="form-control mb-4 flex flex-col gap-2">
    <ProcessResource
      name="CPU"
      defaults={defaultCPULimit}
      bind:value={cpuLimitAmount}
    />
    <ProcessResource
      name="Memory"
      withUnit={true}
      defaults={defaultMemLimit}
      bind:value={memLimitAmount}
      bind:unit={memLimitUnit}
    />
    <ProcessResource
      name="Memory"
      typeLabel="Reserve"
      withUnit={true}
      defaults={defaultMemReservation}
      bind:value={memReservedAmount}
      bind:unit={memReservedUnit}
    />
  </div>

  {#if $setResourcesMutation.isError}
    <Error action="updating resources" error={$setResourcesMutation.error} />
  {/if}

  <button
    class="btn btn-primary"
    class:btn-disabled={$setResourcesMutation.isError}
    class:loading={$setResourcesMutation.isLoading}
    on:click={submit}
  >
    Save
  </button>
</div>
