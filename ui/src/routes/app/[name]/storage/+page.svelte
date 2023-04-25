<script>
  import { page } from "$app/stores";
  import {
    useMutation,
    useQuery,
    useQueryClient,
  } from "@sveltestack/svelte-query";
  import { getAppStorages, mountAppStorage, unmountAppStorage } from "$lib/api";

  import Icon from "$common/Icon.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";

  import StorageMount from "./StorageMount.svelte";
  import CreateStorageModal from "./CreateStorageModal.svelte";
  import Card from "$common/Card.svelte";

  const appName = $page.params.name;

  const queryClient = useQueryClient();
  const queryKey = [{ appName }, "getAppStorages"];
  const storageReport = useQuery(queryKey, () => getAppStorages(appName));
  const invalidateStorages = () => queryClient.invalidateQueries(queryKey);

  let createModalOpen = false;
  const mountMutation = useMutation(
    (options) => mountAppStorage(appName, options),
    {
      onSuccess: () => {
        createModalOpen = false;
        invalidateStorages();
      },
    }
  );

  let loading = {};
  const unmountMutation = useMutation(
    (options) => {
      loading[options["hostDir"]] = true;
      return unmountAppStorage(appName, options);
    },
    {
      onSuccess: invalidateStorages,
      onSettled: (_, __, options) => (loading[options["hostDir"]] = false),
    }
  );
</script>

<QueryDataWrapper query={storageReport} action="fetching app storage">
  <Card title="Storage Mounts">
    <div class="flex flex-col gap-3">
      {#each $storageReport.data["mounts"] as mount, i}
        <StorageMount
          {...mount}
          loading={loading[mount["host_dir"]]}
          on:unmount={({ detail }) => $unmountMutation.mutate(detail)}
        />
      {/each}
    </div>

    <div slot="actions">
      <button class="btn gap-2" on:click={() => (createModalOpen = true)}>
        <Icon type="add" />
        New Storage Mount
      </button>
    </div>
  </Card>
</QueryDataWrapper>

<CreateStorageModal
  bind:open={createModalOpen}
  loading={$mountMutation.isLoading}
  error={$mountMutation.error}
  on:create={({ detail }) => $mountMutation.mutate(detail)}
/>
