<script>
  import { useQuery } from "@sveltestack/svelte-query";
  import { getVersions } from "$lib/api";
  import { appTheme } from "$lib/stores";

  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Card from "$common/Card.svelte";
  import Cards from "$common/Cards.svelte";

  const versions = useQuery("getVersions", getVersions);

  const themes = { auto: "auto", light: "corporate", dark: "business" };
  const themeSelected = (e) => appTheme.set(themes[e.target.value]);
</script>

<Cards>
  <Card title="Theme">
    <div class="p-2">
      <select class="select select-bordered w-fit" on:change={themeSelected}>
        {#each Object.keys(themes) as label}
          <option selected={themes[label] === $appTheme}>{label}</option>
        {/each}
      </select>
    </div>
  </Card>

  <Card title="Versions">
    <QueryDataWrapper query={versions} action="getting version info">
      <div class="p-2 flex flex-col gap-4">
        <div class="stats shadow w-fit">
          <div class="stat">
            <div class="stat-title">Shokku</div>
            <div class="stat-value">{$versions.data["shokku"]}</div>
          </div>
        </div>

        <div class="stats shadow w-fit">
          <div class="stat">
            <div class="stat-title">Dokku</div>
            <div class="stat-value">{$versions.data["dokku"]}</div>
          </div>
        </div>
      </div>
    </QueryDataWrapper>
  </Card>
</Cards>
