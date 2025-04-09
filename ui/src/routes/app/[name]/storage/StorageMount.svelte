<script>
  import Icon from "$common/Icon.svelte";
  import { createEventDispatcher } from "svelte";
  import ConfirmationModal from "$common/ConfirmationModal.svelte";

  export let hostDir;
  export let mountDir;
  export let loading;
  export let isBuildMount;
  export let isRunMount;
  export let isDeployMount;

  let hostDokkuDir;
  let isDokkuManaged = false;

  const dokkuDir = "/var/lib/dokku/data/storage/";
  $: if (hostDir) {
    if (hostDir.startsWith(dokkuDir)) {
      isDokkuManaged = true;
      hostDokkuDir = hostDir.slice(dokkuDir.length);
    }
  }

  let confirmationModalOpen = false;

  const dispatch = createEventDispatcher();

  const tryUnmount = (e) => {
    const restart = e.detail.extraOptionChecked;
    dispatch("unmount", { hostDir, mountDir, restart });
  };
</script>

<div class="flex flex-row gap-2">
  <div class="flex items-center flex-grow">
    <label class="input-group input-group-md">
      <span class="w-auto">{isDokkuManaged ? "Dokku Storage Name" : "Docker Volume Name"}</span>
      <input
        type="text"
        value={isDokkuManaged ? hostDokkuDir : hostDir}
        class="input input-md input-bordered flex-grow"
        disabled
      />
    </label>
  </div>

  <div class="flex-grow">
    <label class="input-group input-group-md">
      <span>Container Path</span>
      <input
        type="text"
        class="input input-md input-bordered w-auto flex-grow"
        disabled
        value={mountDir}
      />
    </label>
  </div>
  <button
    class="btn btn-ghost btn-circle text-error-content ml-2"
    class:loading
    on:click={() => (confirmationModalOpen = true)}
  >
    <Icon type="remove" />
  </button>
</div>

<ConfirmationModal
  name="unmount-storage"
  action="unmount this storage"
  extraOption="Restart app?"
  bind:open={confirmationModalOpen}
  on:accepted={tryUnmount}
/>
