<script>
  import { page } from "$app/stores";
  import {
    getAppsList,
    getServiceLinkedApps,
    linkServiceToApp,
    unlinkServiceFromApp,
  } from "$lib/api";
  import { commandExecutionIds, executionIdDescriptions } from "$lib/stores";
  import Loader from "$common/Loader.svelte";
  import Error from "$common/Error.svelte";

  import LinkCard from "$components/links/LinkCard.svelte";
  import {
    useMutation,
    useQuery,
    useQueryClient,
  } from "@sveltestack/svelte-query";
  import Cards from "$common/Cards.svelte";
  import Card from "$common/Card.svelte";

  export let data = {};

  const serviceType = data.serviceType;
  const serviceName = $page.params.name;

  const queryClient = useQueryClient();
  const linkedQueryKey = [{ serviceName }, "getServiceLinkedApps"];

  const allApps = useQuery("getAppsList", getAppsList);
  const getLinkedApps = useQuery(linkedQueryKey, () =>
    getServiceLinkedApps(serviceName, serviceType)
  );
  const invalidateLinkedAppsQuery = () =>
    queryClient.invalidateQueries(linkedQueryKey);

  let linkedApps = [];
  let unlinkedApps = [];
  $: if ($allApps.isSuccess && $getLinkedApps.isSuccess) {
    linkedApps = $getLinkedApps.data;
    unlinkedApps = $allApps.data.filter((service) => {
      return linkedApps.find((el) => service.name === el.name) === undefined;
    });
  }

  let linking = {};
  const mutationSideEffects = {
    onSuccess: invalidateLinkedAppsQuery,
    onSettled: (_, __, { appName }) => {
      linking[appName] = false;
    },
  };

  const linkAppMutation = useMutation(async ({ appName, options }) => {
    linking[appName] = true;
    const id = await linkServiceToApp(serviceName, appName, options);
    $executionIdDescriptions[id] = `Linking ${serviceName} to ${appName}`;
    return await commandExecutionIds.addID(id);
  }, mutationSideEffects);

  const unlinkAppMutation = useMutation(async ({ appName }) => {
    linking[appName] = true;
    const id = await unlinkServiceFromApp(serviceName, appName);
    $executionIdDescriptions[id] = `Unlinking ${serviceName} from ${appName}`;
    return await commandExecutionIds.addID(id);
  }, mutationSideEffects);

  // TODO: show exposed name for apps
</script>

{#if $allApps.isLoading || $getLinkedApps.isLoading}
  <Loader />
{/if}

{#if $allApps.isError}
  <Error action="loading apps" error={$allApps.error} />
{/if}

{#if $getLinkedApps.isError}
  <Error action="loading linked apps" error={$getLinkedApps.error} />
{/if}

{#if $allApps.isSuccess && $getLinkedApps.isSuccess}
  <Cards>
    {#if linkedApps.length > 0}
      <Card title="Linked Apps">
        <div class="flex flex-row flex-wrap gap-2">
          {#each linkedApps as appName}
            <LinkCard
              {serviceName}
              {serviceType}
              {appName}
              loading={linking[appName]}
              isLinked={true}
              on:unlink={(e) => $unlinkAppMutation.mutate(e.detail)}
            />
          {/each}
        </div>
      </Card>
    {/if}

    <Card title="Unlinked Apps">
      <div class="flex flex-row flex-wrap w-auto gap-2">
        {#each unlinkedApps as appName}
          <LinkCard
            {serviceName}
            {serviceType}
            {appName}
            loading={linking[appName]}
            isLinked={false}
            on:link={(e) => $linkAppMutation.mutate(e.detail)}
          />
        {/each}
      </div>
    </Card>
  </Cards>
{/if}

{#if $linkAppMutation.isError}
  <Error action="linking app" error={$linkAppMutation.error} />
{/if}

{#if $unlinkAppMutation.isError}
  <Error action="unlinking service" error={$unlinkAppMutation.error} />
{/if}
