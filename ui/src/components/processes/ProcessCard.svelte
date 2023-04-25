<script>
  import Icon from "$common/Icon.svelte";

  import ProcessResourceView from "./ProcessResourceView.svelte";
  import ProcessResourceEditor from "./ProcessResourceEditor.svelte";
  import ProcessScaleSelector from "./ProcessScaleSelector.svelte";
  import ProcessDeploymentEditor from "./ProcessDeploymentEditor.svelte";
  import { createEventDispatcher } from "svelte";

  export let processName;
  export let resourceDefaults;

  export let report;
  export let checksReport;

  let resources;
  let scale;

  let cpuLimit;
  let memLimit;
  let memReserved;

  const getResourceByKey = (key, resource) => {
    let resSettings = resources[key][resource];
    if (resSettings === null) return {};
    return {
      amount: resSettings["amount"],
      unit: resSettings["type"]["suffix"],
    };
  };
  const resourceLimit = (resource) => getResourceByKey("limits", resource);
  const resourceReservation = (resource) =>
    getResourceByKey("reservations", resource);

  $: if (report) {
    scale = report["scale"];
    resources = report["resources"];

    cpuLimit = resourceLimit("cpu");
    memLimit = resourceLimit("memory");
    memReserved = resourceReservation("memory");
  }

  const resourceView = 0;
  const resourceEditView = 1;
  const deploymentEditView = 2;

  let currentView = resourceView;
  const setView = (view) => (currentView = view);

  const dispatch = createEventDispatcher();
  const resourcesEdited = () => {
    dispatch("resourcesEdited");
    setView(resourceView);
  };
</script>

<div
  class="w-full bg-neutral text-neutral-content rounded-lg p-4 flex flex-col gap-2"
>
  <div class="flex flex-row items-center gap-2 bg-neutral-focus rounded-lg p-3">
    <div class="">
      <span class="text-xl">{processName}</span>
    </div>
    <div class="flex-grow" />
    <div class="">
      <ProcessScaleSelector {processName} {scale} />
    </div>
  </div>
  <div class="flex flex-row gap-2">
    {#if currentView !== resourceView}
      <button
        on:click={() => setView(resourceView)}
        class="btn btn-sm gap-2 w-52 bg-neutral-focus"
      >
        <Icon type="left" size="sm" />
        back
      </button>
    {:else}
      <div class="">
        <button
          on:click={() => setView(resourceEditView)}
          class="btn btn-sm btn-outline btn-ghost text-neutral-content gap-2"
        >
          Edit Resources
          <Icon type="build" size="sm" />
        </button>
      </div>
      <div class="">
        <button
          on:click={() => setView(deploymentEditView)}
          class="btn btn-sm btn-outline btn-ghost text-neutral-content gap-2"
        >
          Edit Deployment Settings
          <Icon type="cube" size="sm" />
        </button>
      </div>
    {/if}
  </div>

  <div class="rounded-lg">
    {#if report}
      {#if currentView === resourceView}
        <ProcessResourceView {cpuLimit} {memLimit} {memReserved} />
      {:else if currentView === resourceEditView}
        <ProcessResourceEditor
          {processName}
          {resourceDefaults}
          on:successfulEdit={resourcesEdited}
          {cpuLimit}
          {memLimit}
          {memReserved}
        />
      {:else if currentView === deploymentEditView}
        <ProcessDeploymentEditor {processName} report={checksReport} />
      {/if}
    {/if}
  </div>
</div>
