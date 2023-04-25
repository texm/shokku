<script>
  import CloneServiceButton from "./CloneServiceButton.svelte";
  import DestroyServiceButton from "./DestroyServiceButton.svelte";
  import Cards from "$common/Cards.svelte";
  import Card from "$common/Card.svelte";

  export let data = {};
  const serviceType = data.serviceType;
</script>

<Cards>
  <Card title="Clone Service">
    <CloneServiceButton />
  </Card>

  <Card title="Destroy Service">
    <DestroyServiceButton {serviceType} />
  </Card>
</Cards>
