<script>
  import { useQuery } from "@sveltestack/svelte-query";
  import { getDockerRegistryReport } from "$lib/api";

  import Card from "$common/Card.svelte";
  import DockerRegistryButton from "./DockerRegistryButton.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";

  let report = useQuery("getDockerRegistryReport", getDockerRegistryReport);
  let pushOnRelease;
  let server;

  $: if ($report.data) {
    pushOnRelease = $report.data["push_on_release"];
    server = $report.data["server"];
  }
</script>

<Card title="Docker Registry Management">
  <QueryDataWrapper query={report} action="loading registry settings">
    <DockerRegistryButton {server} />

    <div>
      <label class="label cursor-pointer w-52">
        <span class="label-text">Push image on release</span>
        <input
          type="checkbox"
          class="toggle"
          bind:checked={pushOnRelease}
          on:change={() => alert("todo")}
        />
      </label>
    </div>
  </QueryDataWrapper>
</Card>
