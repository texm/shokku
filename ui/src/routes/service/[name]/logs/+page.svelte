<script>
  import { page } from "$app/stores";
  import { getServiceLogs } from "$lib/api";
  import { useQuery } from "@sveltestack/svelte-query";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Logs from "$common/Logs.svelte";
  import Card from "$common/Card.svelte";

  export let data = {};

  const serviceType = data.serviceType;
  const serviceName = $page.params.name;

  const queryKey = [{ serviceName }, "getServiceLogs"];
  const serviceLogs = useQuery(queryKey, () =>
    getServiceLogs(serviceName, serviceType)
  );

  let logs = $serviceLogs.data || [];
</script>

<QueryDataWrapper query={serviceLogs} action="loading logs">
  <Card title="Service Logs">
    {#if logs.length > 0}
      <Logs {logs} />
    {:else}
      <span>No logs available</span>
    {/if}
  </Card>
</QueryDataWrapper>
