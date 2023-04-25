<script>
  import {
    getAppProcesses,
    getAppProcessReport,
    getAppDeployChecksReport,
  } from "$lib/api";
  import Loader from "$common/Loader.svelte";
  import Error from "$common/Error.svelte";
  import ProcessCard from "./ProcessCard.svelte";
  import { page } from "$app/stores";
  import { useQuery, useQueryClient } from "@sveltestack/svelte-query";
  import { createEventDispatcher } from "svelte";

  const appName = $page.params.name;

  const dispatch = createEventDispatcher();

  const queryClient = useQueryClient();

  const psReportQueryKey = ["getAppProcessReport", appName];
  const processReport = useQuery(psReportQueryKey, () =>
    getAppProcessReport(appName)
  );

  const psQueryKey = ["getAppProcesses", appName];
  const appProcesses = useQuery(psQueryKey, () => getAppProcesses(appName));

  const checksQueryKey = ["getAppDeployChecksReport", appName];
  const checksReport = useQuery(checksQueryKey, () =>
    getAppDeployChecksReport(appName)
  );

  const invalidateProcessReport = () =>
    queryClient.invalidateQueries(psReportQueryKey);

  let reports = {};
  let resourceDefaults = {};
  $: if ($processReport.isSuccess) {
    reports = $processReport.data["processes"];
    resourceDefaults = $processReport.data["resource_defaults"];
  }
</script>

<div class="flex flex-col gap-2">
  {#if $processReport.isLoading || $appProcesses.isLoading || $checksReport.isLoading}
    <Loader />
  {:else if $processReport.isError}
    <Error action="loading app process report" error={$processReport.error} />
  {:else if $appProcesses.isError}
    <Error action="loading app processes" error={$appProcesses.error} />
  {:else if $checksReport.isError}
    <Error action="loading deploy checks" error={$checksReport.error} />
  {:else}
    {#each $appProcesses.data as processName, i}
      <ProcessCard
        {processName}
        {resourceDefaults}
        checksReport={$checksReport.data}
        report={reports[processName]}
        on:resourcesEdited={invalidateProcessReport}
      />
    {/each}
  {/if}
</div>
