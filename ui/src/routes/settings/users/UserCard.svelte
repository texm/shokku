<script>
  export let name = "";
  export let source = "";
  export let ssh_keys = [];
</script>

<div class="flex w-fit bg-base-100 p-2 rounded-lg items-center gap-4">
  {#if source === "github"}
    <div class="avatar">
      <div class="w-8 rounded">
        <img
          src="https://avatars.githubusercontent.com/{name}"
          alt="{name} github avatar"
        />
      </div>
    </div>
  {:else}
    <div class="avatar placeholder">
      <div class="bg-neutral-focus text-neutral-content rounded-full w-8">
        <span class="text-lg">{name.substring(0, 1).toUpperCase()}</span>
      </div>
    </div>
  {/if}
  <span class="text-xl">{name}</span>
</div>
