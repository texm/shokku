<script>
  export let name;
  export let value;
  export let defaults;
  export let typeLabel = "Limit";
  export let unit = null;
  export let withUnit = false;

  let enabled = !!value;

  const units = ["b", "k", "m", "g"];
  const onToggled = () => {
    if (!enabled) {
      value = null;
      unit = null;
      return;
    }

    if (!value) value = defaults ? defaults["amount"] : 0;

    if (!unit && withUnit) {
      unit = defaults && defaults["type"] ? defaults["type"]["suffix"] : "m";
    }
  };
</script>

<div class="bg-neutral-focus rounded-lg w-fit p-2">
  <label class="label cursor-pointer w-52">
    <span class="label-text text-neutral-content">{typeLabel} {name}</span>
    <input
      type="checkbox"
      class="toggle"
      bind:checked={enabled}
      on:change={onToggled}
    />
  </label>
  {#if enabled}
    <div>
      <label class="label">
        <span class="label-text text-neutral-content">{name} {typeLabel}</span>
      </label>
      <label class="input-group">
        <span class="text-neutral">{name}</span>
        <input type="number" bind:value class="input input-bordered" />
        {#if withUnit}
          <select class="select select-bordered" bind:value={unit}>
            {#each units as unitOption}
              <option value={unitOption}>{unitOption}</option>
            {/each}
          </select>
        {/if}
      </label>
    </div>
  {/if}
</div>
