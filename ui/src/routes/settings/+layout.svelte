<script>
  import ContentPage from "$common/ContentPage.svelte";
  import Sidebar from "$common/Sidebar.svelte";

  export let data;

  const pages = [
    { name: "Overview", path: "" },
    { name: "Users", path: "users" },
    { name: "Git", path: "git" },
    { name: "Docker", path: "docker" },
    { name: "Domains", path: "domains" },
    { name: "Scheduler", path: "scheduler" },
    { name: "Plugins", path: "plugins" },
  ];
</script>

<ContentPage>
  <div slot="sidebar">
    <Sidebar prefix="/settings" {pages} />
  </div>

  <div slot="content">
    <slot />
  </div>
</ContentPage>
