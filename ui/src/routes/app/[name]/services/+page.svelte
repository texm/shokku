<script>
  import {
    listServices,
    listAppServices,
    linkServiceToApp,
    unlinkServiceFromApp,
  } from "$lib/api";
  import { commandExecutionIds, executionIdDescriptions } from "$lib/stores";
  import { page } from "$app/stores";

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

  const appName = $page.params.name;

  const queryClient = useQueryClient();
  const linkedQueryKey = [{ appName }, "listAppServices"];

  const allServices = useQuery("listServices", listServices);
  const getAppServices = useQuery(linkedQueryKey, () =>
    listAppServices(appName)
  );

  let linkedServices = [];
  let unlinkedServices = [];
  $: if ($getAppServices.isSuccess && $allServices.isSuccess) {
    linkedServices = $getAppServices.data;
    unlinkedServices = $allServices.data.filter((service) => {
      return (
        linkedServices.find((el) => service.name === el.name) === undefined
      );
    });
  }

  const invalidateServicesQuery = () =>
    queryClient.invalidateQueries(linkedQueryKey);

  let loading = {};
  const linkService = async ({ serviceName, options }) => {
    loading[serviceName] = true;
    const id = await linkServiceToApp(serviceName, appName, options);
    $executionIdDescriptions[id] = `Linking ${serviceName} to ${appName}`;
    return await commandExecutionIds.addID(id);
  };
  const unlinkService = async ({ serviceName }) => {
    loading[serviceName] = true;
    const id = await unlinkServiceFromApp(serviceName, appName);
    $executionIdDescriptions[id] = `Unlinking ${serviceName} from ${appName}`;
    return await commandExecutionIds.addID(id);
  };
  const mutationOutcomes = {
    onSuccess: () => invalidateServicesQuery(),
    onSettled: (_, __, { serviceName }) => {
      loading[serviceName] = false;
    },
  };
  const linkMutation = useMutation(linkService, mutationOutcomes);
  const unlinkMutation = useMutation(unlinkService, mutationOutcomes);
</script>

{#if $allServices.isLoading || $getAppServices.isLoading}
  <Loader />
{/if}

{#if $allServices.isError}
  <Error action="loading services" error={$allServices.error} />
{/if}

{#if $getAppServices.isError}
  <Error action="loading linked services" error={$getAppServices.error} />
{/if}

{#if $allServices.isSuccess && $getAppServices.isSuccess}
  <Cards>
    {#if linkedServices.length > 0}
      <Card title="Linked">
        <div class="flex flex-row flex-wrap gap-2">
          {#each linkedServices as service}
            <LinkCard
              cardType="service"
              {appName}
              isLinked={true}
              serviceName={service.name}
              serviceType={service.type}
              on:unlink={({ detail }) => $unlinkMutation.mutate(detail)}
              loading={loading[service.name]}
            />
          {/each}
        </div>
      </Card>
    {/if}

    {#if unlinkedServices.length > 0}
      <Card title="Unlinked">
        <div class="flex flex-row flex-wrap w-auto gap-2">
          <!-- unlinked services -->
          {#each unlinkedServices as service}
            <LinkCard
              cardType="service"
              {appName}
              isLinked={false}
              serviceName={service.name}
              serviceType={service.type}
              on:link={({ detail }) => $linkMutation.mutate(detail)}
              loading={loading[service.name]}
            />
          {/each}
        </div>
      </Card>
    {:else}
      <Card title="No Services">
        <a href="/service/new" class="link">
          <button class="btn btn-secondary">Create a service</button>
        </a>
      </Card>
    {/if}
  </Cards>
{/if}

{#if $linkMutation.isError}
  <Error action="linking service" error={$linkMutation.error} />
{/if}

{#if $unlinkMutation.isError}
  <Error action="unlinking service" error={$unlinkMutation.error} />
{/if}
