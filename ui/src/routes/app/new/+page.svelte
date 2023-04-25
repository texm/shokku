<script>
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query";
  import { createApp } from "$lib/api";
  import { goto } from "$app/navigation";
  import Error from "$common/Error.svelte";
  import Card from "$common/Card.svelte";

  const queryClient = useQueryClient();
  const createAppMutation = useMutation((appName) => createApp(appName), {
    onSuccess: () => queryClient.invalidateQueries("getAppsList"),
  });

  let newAppName = "";
  let creationModalOpen = false;
  const appCreationConfirmed = async () => {
    const success = await $createAppMutation.mutateAsync(newAppName);
    if (!success) return;
    creationModalOpen = false;
    await goto(`/app/${newAppName}/setup`);
  };
</script>

<div class="flex flex-row justify-center p-4">
  <!--div class="hidden md:inline flex-grow" /-->
  <div class="p-3 w-fit">
    <Card title="Create a new app">
      <div class="form-control">
        <label class="input-group">
          <span class="label-text text-base-content">App Name</span>
          <input
            bind:value={newAppName}
            type="text"
            class="input input-bordered"
            disabled={$createAppMutation.isLoading}
          />
        </label>
      </div>

      <div slot="actions">
        <button
          class="btn btn-primary"
          on:click={appCreationConfirmed}
          class:loading={$createAppMutation.isLoading}
        >
          Create
        </button>
      </div>

      {#if $createAppMutation.error}
        <Error action="creating new app" error={$createAppMutation.error} />
      {/if}
    </Card>
    <!--div class="hidden md:inline flex-grow" /-->
  </div>
</div>
