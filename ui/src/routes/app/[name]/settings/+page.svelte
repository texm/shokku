<script>
  import DestroyAppButton from "./DestroyAppButton.svelte";
  import RenameAppButton from "./RenameAppButton.svelte";
  import Cards from "$common/Cards.svelte";
  import Card from "$common/Card.svelte";
</script>

<Cards>
  <Card title="Rename App">
    <RenameAppButton />
  </Card>

  <Card title="Destroy App">
    <DestroyAppButton />
  </Card>
</Cards>
