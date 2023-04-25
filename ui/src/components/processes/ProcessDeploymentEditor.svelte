<script>
  import { setAppProcessDeployChecksState } from "$lib/api";
  import { page } from "$app/stores";

  let appName = $page.params.name;

  export let processName;
  export let report;

  let globalDisabled = false;
  let isDisabled = false;
  let globalSkipped = false;
  let isSkipped = false;

  $: if (report) {
    globalDisabled = report["all_disabled"];
    globalSkipped = report["all_skipped"];
    if (processName in report["disabled_processes"]) isDisabled = true;
    if (processName in report["skipped_processes"]) isSkipped = true;
  }

  let loading = false;
  const stateChanged = async () => {
    loading = true;
    try {
      let state = "enabled";
      if (isSkipped) state = "skipped";
      if (isDisabled) state = "disabled";
      await setAppProcessDeployChecksState(appName, processName, state);
    } finally {
      loading = false;
    }
  };
</script>

<div
  class:tooltip={globalSkipped || isDisabled || globalDisabled}
  class="tooltip-warning"
  data-tip="Checks are {globalSkipped
    ? 'skipped'
    : 'disabled'} for the entire app"
>
  <label class="label cursor-pointer w-96 mt-2">
    <span class="label-text text-neutral-content">Skip Deploy Checks</span>
    <input
      type="checkbox"
      class="toggle"
      disabled={globalDisabled || isDisabled}
      bind:checked={isSkipped}
      on:change={stateChanged}
    />
  </label>
</div>

<div
  class:tooltip={globalDisabled}
  class="tooltip-warning"
  data-tip="Checks are disabled for the entire app"
>
  <label class="label cursor-pointer w-96 mt-2">
    <span class="label-text text-neutral-content"
      >Disable Deploy Checks (may cause downtime)</span
    >
    <input
      type="checkbox"
      class="toggle"
      disabled={globalDisabled}
      bind:checked={isDisabled}
      on:change={stateChanged}
    />
  </label>
</div>
