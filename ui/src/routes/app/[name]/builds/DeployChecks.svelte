<script>
  import { page } from "$app/stores";
  import { useMutation, useQuery } from "@sveltestack/svelte-query";
  import { getAppDeployChecksReport, setAppDeployChecksState } from "$lib/api";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";

  const appName = $page.params.name;

  let allDisabled = false;
  let allSkipped = false;
  const getReport = async () => {
    const report = await getAppDeployChecksReport(appName);
    allDisabled = report["all_disabled"];
    allSkipped = report["all_skipped"];
    return report;
  };

  const queryKey = [{ appName }, "getAppDeployChecks"];
  const checksReport = useQuery(queryKey, getReport);
  const stateMutation = useMutation((state) =>
    setAppDeployChecksState(appName, state)
  );

  const updateSkipped = () => {
    if (allDisabled) return;
    $stateMutation.mutate(allSkipped ? "skipped" : "enabled");
  };

  const updateDisabled = () =>
    $stateMutation.mutate(allDisabled ? "disabled" : "enabled");
</script>

<QueryDataWrapper query={checksReport} action="loading deploy checks state">
  <label class="label mt-2" class:cursor-pointer={!allDisabled}>
    <span class="label-text">Skip Deploy Checks</span>
    <div
      class="tooltip-warning"
      class:tooltip={allDisabled}
      data-tip="Checks are disabled"
    >
      <input
        type="checkbox"
        class="toggle"
        disabled={allDisabled || $checksReport.isLoading}
        bind:checked={allSkipped}
        on:change={updateSkipped}
      />
    </div>
  </label>
  <label class="label cursor-pointer mt-2">
    <span class="label-text">Disable Deploy Checks</span>
    <input
      type="checkbox"
      class="toggle"
      disabled={$checksReport.isLoading}
      bind:checked={allDisabled}
      on:change={updateDisabled}
    />
  </label>
</QueryDataWrapper>
