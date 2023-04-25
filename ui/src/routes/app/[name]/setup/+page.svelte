<script>
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";
  import { useMutation, useQueryClient } from "@sveltestack/svelte-query";
  import { setupApp, setupAppAsync } from "$lib/api";
  import { commandExecutionIds, executionIdDescriptions } from "$lib/stores";

  import Steps from "$common/Steps.svelte";

  import Select from "./steps/Select.svelte";
  import Configure from "./steps/Configure.svelte";
  import Confirm from "./steps/Confirm.svelte";

  import GitNew from "./configs/GitNew.svelte";
  import GitSync from "./configs/GitSync.svelte";
  import DockerImage from "./configs/DockerImage.svelte";

  const appName = $page.params.name;

  const steps = [
    { label: "Select an app source", component: Select },
    { label: "Configure source", component: Configure },
    { label: "Confirm", component: Confirm },
  ];

  const sources = {
    "new-repo": {
      label: "Create New Git Repo",
      createText: "create a git repo",
      component: GitNew,
    },
    "sync-repo": {
      label: "Sync Existing Git Repo",
      createText: "sync an existing git repo",
      component: GitSync,
    },
    "pull-image": {
      label: "Docker Image",
      createText: "pull a docker image",
      component: DockerImage,
    },
    // "upload-archive": {label: "Upload Archive File", component: Archive, options: null},
  };

  const doSyncSetup = ({ source, options }) => {
    return setupApp(appName, source, options);
  };

  const doAsyncSetup = async ({ source, options }) => {
    const id = await setupAppAsync(appName, source, options);
    $executionIdDescriptions[id] = `Setting up ${appName}`;
    const success = await commandExecutionIds.addID(id);
    if (success) return true;
    throw new Error("execution was unsuccessful");
  };

  const queryClient = useQueryClient();

  const onSuccess = async () => {
    await queryClient.invalidateQueries([{ appName }]);
    await goto(`/app/${appName}`);
  };
  const setupAppMutation = useMutation(doSyncSetup, { onSuccess });
  const setupAppAsyncMutation = useMutation(doAsyncSetup, { onSuccess });

  const trySetup = () => {
    const source = data.selectedSource;
    const options = data.sourceOptions[source];
    if (source === "new-repo") $setupAppMutation.mutate({ source, options });
    else $setupAppAsyncMutation.mutate({ source, options });
  };

  const props = { sources };
  let data = {};
  let loading;
  $: loading = $setupAppMutation.isLoading || $setupAppAsyncMutation.isLoading;
</script>

<Steps {steps} {props} {data} {loading} on:complete={trySetup} />
