<script>
  import { page } from "$app/stores";
  import { getAppLogs } from "$lib/api";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Logs from "$common/Logs.svelte";
  import { useQuery } from "@sveltestack/svelte-query";
  import Card from "$common/Card.svelte";

  const appName = $page.params.name;
  const queryKey = [{ appName }, "getAppLogs"];
  const logsQuery = useQuery(queryKey, () => getAppLogs(appName));
</script>

<QueryDataWrapper query={logsQuery} action="fetching logs">
  <Card title="Logs">
    {#if !$logsQuery.data || $logsQuery.data.length === 0}
      <p>No logs available</p>
    {:else}
      <Logs logs={$logsQuery.data} />
    {/if}
  </Card>
</QueryDataWrapper>
