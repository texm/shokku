<script>
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import { getAppOverview } from "$lib/api";
  import { useQuery } from "@sveltestack/svelte-query";

  import Icon from "$common/Icon.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Code from "$common/Code.svelte";
  import Cards from "$common/Cards.svelte";
  import Card from "$common/Card.svelte";

  import ProcessesOverview from "$components/processes/ProcessesOverview.svelte";

  const appName = $page.params.name;

  const appOverviewQueryKey = [{ appName }, "getAppOverview"];
  const appOverview = useQuery(appOverviewQueryKey, () =>
    getAppOverview(appName)
  );

  let gitPushLines = [];
  let isSetup = false;
  let setupMethod = "";
  let isDeployed = false;
  let isRunning = false;
  let numProcesses = 0;

  $: if ($appOverview.isSuccess && $appOverview.data) {
    const data = $appOverview.data;
    isSetup = data["is_setup"];
    setupMethod = data["setup_method"];
    isDeployed = data["is_deployed"];
    isRunning = data["is_running"];
    numProcesses = data["num_processes"];

    let deployBranch = data["git_deploy_branch"];
    gitPushLines = [
      `git remote add dokku dokku@${$page.url.host}:${appName}`,
      `git push dokku ${deployBranch}:master`,
    ];
  }
</script>


<Cards>
  <Card title={isDeployed ? "Processes" : "Almost there..."}>
    <QueryDataWrapper query={appOverview} action="loading app overview">
      {#if isDeployed}
        <ProcessesOverview />
      {:else}
        <div class="hero p-6">
          <div class="hero-content text-center">
            <div class="max-w-md">
              {#if isSetup && setupMethod === "git"}
                <p class="py-2">Add the remote and push your code to deploy:</p>
                <Code lines={gitPushLines} prefix=">" />
              {:else}
                <p class="py-6">Some additional setup required to deploy</p>
                <button
                  class="btn btn-primary gap-2"
                  on:click={() => goto(`/app/${appName}/setup`)}
                >
                  <Icon type="spanner" size="sm" />
                  setup app
                </button>
              {/if}
            </div>
          </div>
        </div>
      {/if}
    </QueryDataWrapper>
  </Card>
</Cards>
