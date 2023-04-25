<script>
  import { page } from "$app/stores";
  import { createEventDispatcher, getContext } from "svelte";
  import Icon from "$common/Icon.svelte";
  import { startService, stopService, restartService } from "$lib/api";
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query";

  export let serviceType;

  const queryClient = useQueryClient();
  const dispatch = createEventDispatcher();

  const serviceInfo = getContext("serviceInfo");
  const serviceName = $page.params.name;

  const serviceMutation = (fn) =>
    useMutation(() => fn(serviceType, serviceName), {
      onSuccess: () => {
        serviceInfo.set({});
        queryClient.invalidateQueries(["getServiceInfo", serviceName]);
      },
    });
  const stopMutation = serviceMutation(stopService);
  const startMutation = serviceMutation(startService);
  const restartMutation = serviceMutation(restartService);

  let loading;
  $: loading =
    $stopMutation.isLoading ||
    $startMutation.isLoading ||
    $restartMutation.isLoading;
</script>

<div class="flex flex-row w-full h-16 bg-base-200 p-4 items-center rounded-lg">
  <div class="flex items-center gap-4">
    {#if serviceType}
      <Icon size="lg" type={serviceType} />
    {/if}
    <span class="text-3xl align-center">
      {serviceName}
    </span>
  </div>
  <div class="flex-grow" />
  <div class="flex flex-row gap-2" class:hidden={$serviceInfo.status === ""}>
    {#if $serviceInfo.status === "running"}
      <button
        class="btn btn-outline btn-sm"
        class:gap-2={!loading}
        class:loading
        on:click={$stopMutation.mutate}
      >
        {#if !loading}
          <Icon type="stop" size="sm" />
        {/if}
        stop
      </button>
      <button
        class="btn btn-outline btn-sm"
        class:gap-2={!loading}
        class:loading
        on:click={$restartMutation.mutate}
      >
        {#if !loading}
          <Icon type="restart" size="sm" />
        {/if}
        restart
      </button>
    {:else}
      <button
        class="btn btn-outline btn-sm"
        class:gap-2={!loading}
        class:loading
        on:click={$startMutation.mutate}
      >
        {#if !loading}
          <Icon type="start" size="sm" />
        {/if}
        start
      </button>
    {/if}
  </div>
</div>
