<script>
  import {
    getNetworksList,
    getAppNetworksReport,
    updateAppNetworks,
  } from "$lib/api";
  import { page } from "$app/stores";
  import {
    useMutation,
    useQuery,
    useQueryClient,
  } from "@sveltestack/svelte-query";
  import Loader from "$common/Loader.svelte";
  import Error from "$common/Error.svelte";
  import AppNetworks from "./AppNetworks.svelte";

  const appName = $page.params.name;

  const queryClient = useQueryClient();
  const reportQueryKey = [{ appName }, "getAppNetworkReport"];

  const networkReport = useQuery(reportQueryKey, () =>
    getAppNetworksReport(appName)
  );
  const networkList = useQuery("getNetworkList", getNetworksList);
  const updateNetworks = useMutation(
    (config) => updateAppNetworks(appName, config),
    { onSuccess: () => queryClient.invalidateQueries(reportQueryKey) }
  );

  let isLoaded;
  $: isLoaded = $networkReport.isSuccess && $networkList.isSuccess;
</script>

{#if $networkReport.isLoading || $networkList.isLoading}
  <Loader />
{/if}

{#if $networkList.isError}
  <Error action="fetching network info" error={$networkList.error} />
{/if}

{#if $networkReport.isError}
  <Error action="fetching app network info" error={$networkReport.error} />
{/if}

{#if isLoaded}
  <AppNetworks
    report={$networkReport.data}
    networks={$networkList.data}
    loading={$updateNetworks.isLoading}
    on:saved={({ detail }) => $updateNetworks.mutate(detail)}
  />
{/if}

{#if $updateNetworks.isError}
  <Error action="updating app network details" error={$updateNetworks.error} />
{/if}
