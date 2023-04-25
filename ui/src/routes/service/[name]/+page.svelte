<script>
  import { page } from "$app/stores";
  import { getContext } from "svelte";
  import Card from "$common/Card.svelte";

  export let data;

  const serviceName = $page.params.name;
  const serviceInfo = getContext("serviceInfo");
</script>

<Card title="Overview">
  <div class="overflow-x-auto">
    <table class="table">
      <tbody>
        {#each Object.keys($serviceInfo) as key}
          <tr>
            <td>{key}</td>
            <td>{$serviceInfo[key]}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</Card>
