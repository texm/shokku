<script>
  import {
    getAppDomainsReport,
    restartApp,
    rebuildApp,
    startApp,
    stopApp,
  } from "$lib/api";

  import { commandExecutionIds, executionIdDescriptions } from "$lib/stores";

  import { page } from "$app/stores";
  import {
    useMutation,
    useQuery,
    useQueryClient,
  } from "@sveltestack/svelte-query";
  import { createEventDispatcher } from "svelte";
  import AppHeaderIconButton from "$src/routes/app/[name]/AppHeaderIconButton.svelte";
  import Icon from "$common/Icon.svelte";

  export let isSetup;
  export let setupMethod;
  export let isRunning;
  export let isDeployed;

  const appName = $page.params.name;

  const queryClient = useQueryClient();
  const queryKey = [{ appName }, "getAppDomainsReport"];
  const domainsReport = useQuery(queryKey, () => getAppDomainsReport(appName));

  const dispatch = createEventDispatcher();
  const onSuccess = () => {
    dispatch("statusChanged");
  };
  const newMutation = (action, fn) => {
    const mutationFn = async () => {
      const id = await fn(appName);
      $executionIdDescriptions[id] = `${action} ${appName}`;
      return commandExecutionIds.addID(id);
    };
    return useMutation(mutationFn, { onSuccess });
  };

  const restartMutation = newMutation("Restart", restartApp);
  const rebuildMutation = newMutation("Rebuild", rebuildApp);
  const startMutation = newMutation("Start", startApp);
  const stopMutation = newMutation("Stop", stopApp);

  let loading;
  $: loading =
    $startMutation.isLoading ||
    $restartMutation.isLoading ||
    $rebuildMutation.isLoading ||
    $stopMutation.isLoading;

  let appDomain;
  $: if ($domainsReport.data) {
    appDomain = null;
    let domains = $domainsReport.data["domains"];
    if (domains && domains.length > 0) {
      appDomain = domains[0];
    }
  }
</script>

<div class="w-full min-h-16 flex flex-row bg-base-200 p-4 items-center rounded-lg">
  <div class="flex flex-row items-center">
    {#if appDomain}
      <a href="http://{appDomain}" class="link link-hover" target="_blank">
        <div class="flex flex-row">
          <span class="text-3xl mr-1">
            {appName}
          </span>
          <Icon type="external-link" size="sm" />
        </div>
      </a>
    {:else}
      <span class="text-3xl align-center">{appName}</span>
    {/if}
  </div>
  <div class="flex-grow" ></div>
  <div class="flex md:flex-row flex-col gap-2">
    {#if isSetup && isDeployed && isRunning}
      <AppHeaderIconButton
        action="rebuild"
        {loading}
        on:clicked={$rebuildMutation.mutate}
      />
    {/if}
    {#if isDeployed && !isRunning}
      <AppHeaderIconButton
        action="start"
        {loading}
        on:clicked={$startMutation.mutate}
      />
    {:else if isRunning}
      <AppHeaderIconButton
        action="stop"
        {loading}
        on:clicked={$stopMutation.mutate}
      />
      <AppHeaderIconButton
        action="restart"
        {loading}
        on:clicked={$restartMutation.mutate}
      />
    {/if}
  </div>
</div>
