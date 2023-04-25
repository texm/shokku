<script>
  import { page } from "$app/stores";
  import { getAppOverview } from "$lib/api";
  import { useQuery, useQueryClient } from "@sveltestack/svelte-query";

  import ContentPage from "$common/ContentPage.svelte";
  import Sidebar from "$common/Sidebar.svelte";
  import AppHeader from "./AppHeader.svelte";

  const appName = $page.params.name;

  let onSetupPage;
  let currentPage;
  $: if ($page.url) {
    onSetupPage = $page.url.pathname.startsWith(`/app/${appName}/setup`);
    currentPage = $page.url.pathname.substring(`/app/${appName}`.length + 1);
  }

  const pages = [
    { name: "Overview", path: "" },
    { name: "Builds", path: "builds" },
    { name: "Domains", path: "domains" },
    { name: "Environment", path: "environment" },
    { name: "Logs", path: "logs" },
    { name: "Network", path: "network" },
    { name: "Services", path: "services" },
    // {"name": "Terminal", "path": "terminal"},
    { name: "Storage", path: "storage" },
    { name: "Settings", path: "settings" },
  ];

  const queryClient = useQueryClient();
  const appOverviewQueryKey = [{ appName }, "getAppOverview"];
  const appOverview = useQuery(appOverviewQueryKey, () =>
    getAppOverview(appName)
  );
  const statusChanged = () =>
    queryClient.invalidateQueries(appOverviewQueryKey);
</script>

<ContentPage>
  <div slot="sidebar" class:hidden={onSetupPage}>
    <Sidebar {pages} prefix={"/app/" + $page.params.name} />
  </div>

  <div slot="header" class:hidden={onSetupPage} class="mb-2">
    {#if $appOverview.isSuccess}
      <AppHeader
        on:statusChanged={statusChanged}
        isDeployed={$appOverview.data["is_deployed"]}
        isRunning={$appOverview.data["is_running"]}
        isSetup={$appOverview.data["is_setup"]}
        setupMethod={$appOverview.data["setup_method"]}
      />
    {/if}
  </div>

  <div slot="content">
    <slot />
  </div>
</ContentPage>
