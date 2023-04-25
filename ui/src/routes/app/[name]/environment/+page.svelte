<script>
  import { page } from "$app/stores";
  import { getAppConfig, setAppConfig } from "$lib/api";
  import { useQuery, useMutation } from "@sveltestack/svelte-query";

  import Error from "$common/Error.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import KVEditor from "$common/KVEditor.svelte";
  import Card from "$common/Card.svelte";

  const appName = $page.params.name;
  const queryKey = [{ appName }, "getAppConfig"];
  const varsQuery = useQuery(queryKey, () => getAppConfig(appName));

  let stateDirty = false;
  const setConfigMutation = useMutation(setAppConfig, {
    onSuccess: () => {
      stateDirty = false;
    },
  });

  const saveVars = async ({ detail }) => {
    let newConfig = {};
    for (let [k, v] of detail) newConfig[k] = v;
    $setConfigMutation.mutate({ appName, newConfig });
  };
</script>

<QueryDataWrapper query={varsQuery} action="getting environment variables">
  <Card title="Environment Variables">
    <KVEditor
      vars={$varsQuery.data}
      saving={$setConfigMutation.isLoading}
      showSaveButton={true}
      neutralButtons={true}
      bind:stateDirty
      on:save={saveVars}
    />
  </Card>
</QueryDataWrapper>

{#if $setConfigMutation.isError}
  <Error
    action="updating environment variables"
    error={$setConfigMutation.error}
  />
{/if}
