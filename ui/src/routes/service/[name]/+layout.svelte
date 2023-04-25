<script>
  import { setContext } from "svelte";
  import { writable } from "svelte/store";
  import { useQuery, useQueryClient } from "@sveltestack/svelte-query";
  import { page } from "$app/stores";
  import { getServiceInfo } from "$lib/api";

  import ContentPage from "$common/ContentPage.svelte";
  import Sidebar from "$common/Sidebar.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";

  import ServiceHeader from "./ServiceHeader.svelte";

  export let data = {};

  const serviceName = data.serviceName;
  $: if (serviceName !== $page.params.name) location.reload();

  const serviceType = data.serviceType;

  let currentPage;
  $: if ($page.url) {
    currentPage = $page.url.pathname.substring(
      `/service/${serviceName}`.length + 1
    );
  }

  const defaultPages = [
    { name: "Overview", path: "" },
    { name: "Apps", path: "apps" },
    { name: "Logs", path: "logs" },
    { name: "Backups", path: "backups" },
    { name: "Settings", path: "settings" },
  ];

  let pages = [...defaultPages];

  const servicePages = {};
  if (serviceType in servicePages) {
    pages = [...defaultPages, ...servicePages[serviceType]];
  }

  const serviceInfo = writable({});
  setContext("serviceInfo", serviceInfo);

  const queryClient = useQueryClient();
  const queryKey = ["getServiceInfo", serviceName];
  const fetchInfo = useQuery(
    queryKey,
    () => getServiceInfo(serviceName, serviceType),
    { onSuccess: serviceInfo.set }
  );
</script>

<ContentPage>
  <div slot="sidebar">
    <Sidebar {pages} prefix="/service/{serviceName}" />
  </div>

  <div slot="header" class="mb-2">
    <ServiceHeader {serviceType} />
  </div>

  <div slot="content" class="max-w-full">
    <QueryDataWrapper query={fetchInfo} action="loading service info">
      <slot />
    </QueryDataWrapper>
  </div>
</ContentPage>
