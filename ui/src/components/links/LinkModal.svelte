<script>
  import { createEventDispatcher } from "svelte";

  import Modal from "$common/Modal.svelte";

  import Generic from "./link-configs/Generic.svelte";

  export let open = false;
  export let preventClose = false;
  export let serviceType = "";

  const typeComponents = {};

  let configComponent = typeComponents[serviceType] || Generic;

  const serviceConfigComponents = {};
  $: if (serviceType && serviceType in serviceConfigComponents)
    configComponent = serviceConfigComponents[serviceType];

  let serviceConfig = {};

  const dispatch = createEventDispatcher();
  const dispatchLinkService = () => {
    dispatch("link", serviceConfig);
    open = false;
  };
</script>

<Modal name="link-app-service" title="Linking" bind:open {preventClose}>
  <span class="text-lg">This will cause your app to restart!</span>

  <div
    class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box"
  >
    <input type="checkbox" />
    <div class="collapse-title text-md font-medium">Advanced Options</div>
    <div class="collapse-content">
      <svelte:component this={configComponent} bind:config={serviceConfig} />
    </div>
  </div>

  <div class="mt-3">
    <button
      class="btn"
      class:loading={preventClose}
      on:click={dispatchLinkService}
    >
      Submit
    </button>
  </div>
</Modal>
