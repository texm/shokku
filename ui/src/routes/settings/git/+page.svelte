<script>
  import Card from "$common/Card.svelte";
  import GitAuthButton from "./GitAuthButton.svelte";
  import GitRemoveAuthButton from "./GitRemoveAuthButton.svelte";
</script>

<Card title="Git Settings">
  <GitAuthButton />

  <GitRemoveAuthButton />
</Card>
