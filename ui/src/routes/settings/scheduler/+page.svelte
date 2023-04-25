<script>
  import Card from "$common/Card.svelte";
  import Cards from "$common/Cards.svelte";
  import DockerLocalSettings from "./DockerLocalSettings.svelte";

  const schedulers = [
    { name: "Docker Local", component: DockerLocalSettings },
    { name: "Kubernetes" },
    { name: "Nomad" },
  ];
  const selected = 0;
</script>

<Cards>
  <Card title="Selected Scheduler">
    <select class="select select-bordered w-fit" disabled>
      {#each schedulers as scheduler, i}
        <option selected={selected === i}>
          {scheduler.name}
        </option>
      {/each}
    </select>
  </Card>

  <!--Card title="{schedulers[selected].name}">
		<svelte:component this={schedulers[selected].component} />
	</Card-->
</Cards>
