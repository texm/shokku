<script>
  // TODO: delete this?

  import { useMutation } from "@sveltestack/svelte-query";
  import { setAppProcessDeployChecksState } from "$lib/api";
  import { page } from "$app/stores";
  import Error from "$common/Error.svelte";

  export let process;
  export let checks;

  const appName = $page.params.name;
  const stateMutation = useMutation((state) =>
    setAppProcessDeployChecksState(appName, process, state)
  );

  const onToggleEnabled = async (e) =>
    $stateMutation.mutate(e.target.checked ? "enabled" : "disabled");

  const onToggleSkipped = async (e) =>
    $stateMutation.mutate(e.target.checked ? "skipped" : "enabled");

  const manualChecksTrigger = async () => {
    alert("todo");
  };
</script>

<span class="text-lg font-bold">{process}</span>

<div class="form-control">
  <label class="label cursor-pointer">
    <span class="label-text">Enable Deployment Checks</span>
    <input
      type="checkbox"
      class="toggle"
      class:disabled={$stateMutation.isLoading}
      on:change={onToggleEnabled}
      checked={checks["enabled"]}
    />
  </label>
</div>

{#if checks["enabled"]}
  <div class="form-control">
    <label class="label cursor-pointer">
      <span class="label-text">Skip Deployment Checks</span>
      <input
        type="checkbox"
        class="toggle"
        class:disabled={$stateMutation.isLoading}
        on:change={onToggleSkipped}
        checked={checks["skipped"]}
      />
    </label>
  </div>

  <button class="btn btn-sm" on:click={manualChecksTrigger}>Run Checks</button>
{/if}

{#if $stateMutation.isError}
  <Error error={$stateMutation.error} />
{/if}
