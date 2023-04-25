<script>
  import DefaultServiceConfig from "../configs/DefaultServiceConfig.svelte";
  import LitestreamConfig from "../configs/LitestreamConfig.svelte";
  import { createEventDispatcher } from "svelte";

  export let props = {};
  export let data = {};

  let serviceType = data.selectedService?.type;
  if (!data.config) data.config = {};
  if (serviceType && !data.config[serviceType]) data.config[serviceType] = {};

  const dispatch = createEventDispatcher();
  const configChanged = ({ detail }) => {
    let { key, enabled, value } = detail;
    if (enabled) data.config[serviceType][key] = value;
    else delete data.config[serviceType][key];

    dispatch("statusChange", { complete: !!data.config[serviceType]["name"] });
  };

  const uniqueConfigs = { sqlite: LitestreamConfig };
  let creationComponent = DefaultServiceConfig;
  $: if (serviceType && serviceType in uniqueConfigs) {
    creationComponent = uniqueConfigs[serviceType];
  }

  let name = data.config[serviceType]["name"];
  $: configChanged({ detail: { key: "name", enabled: true, value: name } });
</script>

<div class="max-h-full pr-4">
  <div
    class="mb-2 p-2 rounded-lg bg-neutral"
    class:border-4={!name}
    class:border-warning={!name}
  >
    <label class="label">
      <span class="label-text text-neutral-content">
        Give this service a unique name:
      </span>
    </label>
    <label class="input-group w-full flex">
      <span class="w-fit">Name</span>
      <input bind:value={name} class="input input-bordered flex-grow" />
    </label>
  </div>

  <svelte:component
    this={creationComponent}
    config={data.config[serviceType]}
    on:changed={configChanged}
  />
</div>
