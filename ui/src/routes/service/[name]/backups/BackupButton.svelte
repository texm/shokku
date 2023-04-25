<script>
  import { page } from "$app/stores";
  import { doServiceBackup } from "$lib/api";
  import Error from "$common/Error.svelte";
  import { commandExecutionIds, executionIdDescriptions } from "$lib/stores";
  import { useMutation } from "@sveltestack/svelte-query";

  export let ready = false;

  const serviceName = $page.params.name;

  const doBackupMutation = useMutation(async () => {
    const execID = await doServiceBackup(serviceName);
    $executionIdDescriptions[execID] = `Running backup for ${serviceName}`;
    return await commandExecutionIds.addID(execID);
  });
</script>

<button
  class="btn btn-primary w-fit text-left"
  on:click={$doBackupMutation.mutate}
  class:btn-disabled={!ready || $doBackupMutation.isLoading}
>
  Backup Now
</button>

{#if $doBackupMutation.isError}
  <div class="mt-2">
    <Error action="creating backup" error={$doBackupMutation.error} />
  </div>
{/if}
