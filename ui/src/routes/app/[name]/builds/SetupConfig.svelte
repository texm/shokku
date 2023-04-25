<script>
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import { page } from "$app/stores";
  import { useQuery } from "@sveltestack/svelte-query";
  import { getAppSetupConfig } from "$lib/api";

  const appName = $page.params.name;

  const queryKey = [{ appName }, "getAppSetupConfig"];
  const query = useQuery(queryKey, () => getAppSetupConfig(appName));

  let isSetup, method, deploymentBranch, repoUrl, repoGitRef, image;
  $: if ($query.isFetched && $query.data) {
    isSetup = $query.data["is_setup"];
    method = $query.data["method"];
    image = $query.data["image"];
    deploymentBranch = $query.data["deployment_branch"];
    repoUrl = $query.data["repo_url"];
    repoGitRef = $query.data["repo_git_ref"];
  }
</script>

<QueryDataWrapper {query} action="getting setup config">
  {#if isSetup}
    {#if method === "git"}
      <span
        >Building on git push to branch <strong>{deploymentBranch}</strong
        ></span
      >
    {:else if method === "sync-repo"}
      <span
        >Synced from remote git repo <strong>{repoUrl}</strong>
        using git ref {repoGitRef}</span
      >
    {:else if method === "docker"}
      <span>Built from image <strong>{image}</strong></span>
    {/if}

    <div class="">
      <a href={`/app/${appName}/setup`}>
        <button class="btn btn-primary"> Update Configuration </button>
      </a>
    </div>
  {:else}
    <div>
      <span class="text-lg text-warning"
        >App needs to be setup before it can be built</span
      >
    </div>
    <a href={`/app/${appName}/setup`}>
      <button class="btn btn-info"> Go to setup </button>
    </a>
  {/if}
</QueryDataWrapper>
