<script>
  import Cards from "$common/Cards.svelte";
  import Card from "$common/Card.svelte";

  import BuilderSelection from "./BuilderSelection.svelte";
  import BuildDirectory from "./BuildDirectory.svelte";
  import DeployChecks from "./DeployChecks.svelte";
  import SetupConfig from "./SetupConfig.svelte";
</script>

<Cards>
  <Card title="Config">
    <div class=""><SetupConfig /></div>
  </Card>

  <Card title="Builder Selection">
    <div class="max-w-xs"><BuilderSelection /></div>
  </Card>

  <Card title="Build Directory">
    <div class="max-w-xs"><BuildDirectory /></div>
  </Card>

  <Card title="Zero-Downtime Deployment Checks">
    <div class="max-w-xs"><DeployChecks /></div>
  </Card>
</Cards>
