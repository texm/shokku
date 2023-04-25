<script>
  import "$src/app.css";
  import { onMount, onDestroy } from "svelte";
  import { startAuthRefresh, stopAuthRefresh } from "$lib/auth";
  import { commandExecutionIds, appTheme } from "$lib/stores";

  import Header from "$components/Header.svelte";
  import { QueryClient, QueryClientProvider } from "@sveltestack/svelte-query";

  import CommandExecutionWindow from "$components/commands/CommandExecutionWindow.svelte";
  import { goto } from "$app/navigation";

  export let data = {};

  let isAuthenticated;
  $: isAuthenticated = !!data.authDetails;

  let executionWindowVisible = null;
  let executionWindowWatchingCompleted = false;
  let showExecutionButton = false;

  $: if ($commandExecutionIds !== null) {
    showExecutionButton =
      $commandExecutionIds.length > 0 || executionWindowWatchingCompleted;
    executionWindowVisible = showExecutionButton;
  }

  const toggleExecutionWindow = () => {
    executionWindowVisible = !executionWindowVisible;
  };

  onMount(startAuthRefresh);
  onDestroy(stopAuthRefresh);

  const checkQueryError = (err) => {
    err = JSON.parse(err.message);
    if (err.message === "setup key invalid") goto("/?invalidate_setup=1");
  };
  const queryClient = new QueryClient();
  queryClient.setDefaultOptions({
    queries: {
      refetchOnMount: false,
      refetchOnWindowFocus: false,
      retry: 1,
      onError: checkQueryError,
    },
  });
</script>

<svelte:head>
  <title>shokku</title>
</svelte:head>

<div data-theme={$appTheme}>
  <QueryClientProvider client={queryClient}>
    <div class="flex flex-col h-screen w-screen max-w-screen max-h-screen">
      <div>
        <Header
          {isAuthenticated}
          {showExecutionButton}
          {executionWindowVisible}
          on:executionButtonClicked={toggleExecutionWindow}
        />
      </div>

      <div class="flex flex-grow min-h-0 h-full">
        <div class="flex-grow w-full min-h-0 h-full max-h-screen">
          <slot />
        </div>
      </div>
    </div>

    <div
      class="absolute z-50 top-24 px-8 w-full"
      class:hidden={!executionWindowVisible}
    >
      <CommandExecutionWindow
        bind:watchingCompleted={executionWindowWatchingCompleted}
      />
    </div>
  </QueryClientProvider>
</div>
