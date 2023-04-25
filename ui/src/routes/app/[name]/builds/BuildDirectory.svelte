<script>
  import {
    getAppBuildDirectory,
    setAppBuildDirectory,
    clearAppBuildDirectory,
  } from "$lib/api";
  import { page } from "$app/stores";
  import {
    useMutation,
    useQuery,
    useQueryClient,
  } from "@sveltestack/svelte-query";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Error from "$common/Error.svelte";

  const appName = $page.params.name;

  const queryClient = useQueryClient();
  const queryKey = [{ appName }, "getAppBuildDirectory"];

  let buildDir = "";
  let enableBuildDirectory = false;

  const onBuildDirFetched = (data) => {
    buildDir = data;
    enableBuildDirectory = !!buildDir;
  };

  const getBuildDir = useQuery(queryKey, () => getAppBuildDirectory(appName), {
    onSuccess: onBuildDirFetched,
  });
  const setDir = useMutation(
    (newDir) => setAppBuildDirectory(appName, newDir),
    {
      onSuccess: (_, newDir) => queryClient.setQueryData(queryKey, newDir),
    }
  );
  const clearDir = useMutation(() => clearAppBuildDirectory(appName), {
    onSuccess: () => queryClient.setQueryData(queryKey, null),
  });

  const onEnableToggled = async (e) => {
    enableBuildDirectory = e.target.checked;
    if (!e.target.checked) $clearDir.mutate(appName);
  };
</script>

<QueryDataWrapper query={getBuildDir} action="loading build directory">
  <label class="label cursor-pointer">
    <span class="label-text">Use custom directory</span>
    <input
      type="checkbox"
      class="toggle"
      disabled={$setDir.isLoading || $clearDir.isLoading}
      checked={enableBuildDirectory}
      on:change={onEnableToggled}
    />
  </label>
  {#if enableBuildDirectory}
    <label class="input-group input-group-md">
      <span class="w-auto">Directory</span>
      <input
        type="text"
        class="input input-md input-bordered flex-grow"
        disabled={$setDir.isLoading || $clearDir.isLoading}
        bind:value={buildDir}
      />
    </label>
    <div class="mt-4">
      <button
        class="btn btn-primary"
        class:loading={$setDir.isLoading || $clearDir.isLoading}
        class:hidden={$getBuildDir.data === buildDir}
        on:click={() => $setDir.mutate(buildDir)}
      >
        Save
      </button>
    </div>
  {/if}
</QueryDataWrapper>

{#if $setDir.isError}
  <Error action="setting build directory" error={$setDir.error} />
{/if}

{#if $clearDir.isError}
  <Error action="clearing build directory" error={$clearDir.error} />
{/if}
