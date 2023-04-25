<script>
  import Icon from "$common/Icon.svelte";
  import { createEventDispatcher } from "svelte";

  export let props = {};
  export let data = {};

  const serviceTypes = [
    { name: "PostgreSQL", type: "postgres" },
    { name: "MySQL", type: "mysql" },
    { name: "MongoDB", type: "mongo" },
    { name: "Redis", type: "redis" },
    // {name: "Litestream (SQLite)", type: "sqlite"},
  ];

  const dispatch = createEventDispatcher();
  const serviceSelected = (svc) => {
    dispatch("statusChange", { complete: true });
    data.selectedService = svc;
  };

  let selectedServiceName = data.selectedService?.name;
  $: selectedServiceName = data.selectedService?.name;
</script>

<span class="text-2xl">Select a service to create:</span>

<div class="flex flex-row flex-wrap gap-2 p-4">
  {#each serviceTypes as svc}
    <div
      class="card w-fit shadow-lg cursor-pointer hover:bg-primary-focus hover:text-primary-content"
      class:bg-base-100={selectedServiceName !== svc.name}
      class:bg-primary={selectedServiceName === svc.name}
      class:text-primary-content={selectedServiceName === svc.name}
    >
      <div
        on:click={() => serviceSelected(svc)}
        class="card-body flex flex-row gap-4"
      >
        <Icon type={svc.type} size="lg" />
        <h2 class="card-title">{svc.name}</h2>
      </div>
    </div>
  {/each}
</div>
