<script>
  import { createEventDispatcher, onMount } from "svelte";
  import Error from "$common/Error.svelte";

  export let props = {};
  export let data = {};

  let { selectedService, config } = data;
  let serviceType = selectedService.type;
  let serviceConfig = config[serviceType];
  let serviceName = serviceConfig.name;
  let configKeys = Object.keys(serviceConfig).filter((key) => key !== "name");

  const dispatch = createEventDispatcher();
  onMount(() => dispatch("statusChange", { complete: true }));
</script>

<span class="text-lg">
  Creating a <strong>{serviceType}</strong> service named
  <strong>{serviceName}</strong>
  with
  {Object.keys(configKeys).length > 0
    ? "the following configuration:"
    : "the default configuration."}
</span>
{#each configKeys as key}
  <div><span class="text-lg">{key}: {serviceConfig[key]}</span></div>
{/each}
