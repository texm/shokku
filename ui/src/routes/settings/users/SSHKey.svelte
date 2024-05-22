<script>
  import Icon from "$common/Icon.svelte";

  export let name = "";
  export let fingerprint = "";

  let tooltipText = "copy";

  const copyFingerprint = (e) => {
    navigator.clipboard.writeText(fingerprint);
    tooltipText = "copied!";
    setTimeout(() => (tooltipText = "copy"), 1000);
  };

  const deleteKey = () => {
    alert("todo");
  };
</script>

<div
  class="flex gap-2 max-w-full overflow-auto text-center items-center bg-base-100 p-4 rounded-lg"
>
  <span class="text-lg">{name}</span>
  <div class="tooltip tooltip-left flex-grow" data-tip={tooltipText}>
    <button class="btn btn-ghost" on:click={copyFingerprint}>
      {fingerprint}
    </button>
  </div>
  <div class="" class:hidden={name === "admin"}>
    <button class="btn btn-square btn-outline btn-sm" on:click={deleteKey}>
      <Icon size="sm" type="delete" />
    </button>
  </div>
</div>
