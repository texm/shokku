<script>
  import { getSelectedBuilder, setSelectedBuilder } from "$lib/api";
  import {
    useMutation,
    useQuery,
    useQueryClient,
  } from "@sveltestack/svelte-query";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Error from "$common/Error.svelte";
  import { page } from "$app/stores";

  const appName = $page.params.name;

  const queryClient = useQueryClient();
  const queryKey = [{ appName }, "getAppBuilder"];
  const selectedBuilder = useQuery(queryKey, () => getSelectedBuilder(appName));
  const updateBuilder = useMutation(
    (newBuilder) => setSelectedBuilder(appName, newBuilder),
    {
      onSuccess: (_, newBuilder) =>
        queryClient.setQueryData(queryKey, newBuilder),
    }
  );

  const builders = {
    auto: "Auto Detect",
    dockerfile: "Dockerfile",
    herokuish: "Herokuish",
    lambda: "Lambda",
    // "pack": "Cloud Native Buildpacks",
    null: "Null Builder",
  };
</script>

<QueryDataWrapper query={selectedBuilder} action="loading selected builder">
  <select
    class="select w-full select-bordered"
    value={$selectedBuilder.data}
    on:change={(e) => $updateBuilder.mutate(e.target.value)}
  >
    {#each Object.keys(builders) as builder}
      <option value={builder}>{builders[builder]}</option>
    {/each}
  </select>
</QueryDataWrapper>

{#if $updateBuilder.isError}
  <Error action="changing selected builder" error={$updateBuilder.error} />
{/if}
