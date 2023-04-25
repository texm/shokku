<script>
  import { useQuery } from "@sveltestack/svelte-query";
  import { getAllAppsOverview, listServices } from "$lib/api";

  import DashboardCardList from "$components/dashboard/DashboardCardList.svelte";

  const queryOpts = {
    refetchOnMount: true,
  };
  const appsOverview = useQuery(
    "getAllAppsOverview",
    getAllAppsOverview,
    queryOpts
  );
  const servicesList = useQuery("listServices", listServices, queryOpts);
  const appStates = {};
  const serviceStates = {};
</script>

<div class="flex flex-col md:flex-row gap-8 p-4 h-full overflow-y-scroll">
  <div class="hidden md:inline flex-grow" />
  <div class="max-w-full md:max-w-lg md:flex-grow">
    <DashboardCardList
      contentType="app"
      query={appsOverview}
      states={appStates}
    />
  </div>
  <div class="max-w-full md:max-w-lg md:flex-grow">
    <DashboardCardList
      contentType="service"
      query={servicesList}
      states={serviceStates}
    />
  </div>
  <div class="hidden md:inline flex-grow" />
</div>
