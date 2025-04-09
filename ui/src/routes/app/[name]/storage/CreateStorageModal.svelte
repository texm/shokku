<script>
  import { createEventDispatcher } from "svelte";

  import Modal from "$common/Modal.svelte";
  import Error from "$common/Error.svelte";

  export let open;
  export let loading;
  export let error;

  const dokkuDir = "/var/lib/dokku/data/storage/";
  const storageTypes = {
    "Docker Volume": { label: "Volume Name" },
    "Dokku Storage": { label: dokkuDir },
  };

  let selectedType = "Dokku Storage";
  let hostDir = "";
  let mountDir = "";
  let typeOptions = Object.keys(storageTypes);

  const dispatch = createEventDispatcher();

  const dispatchCreateStorage = () => {
    const fullHostDir = selectedType === "Dokku Storage" ? `${dokkuDir}${hostDir}` : hostDir;
    const options = { hostDir: fullHostDir, mountDir };
    dispatch("create", options);
  };
</script>

<Modal
  name="create-storage"
  title="Create App Storage"
  bind:open
  preventClose={loading}
>
  <div class="form-control" class:hidden={typeOptions.length < 2}>
    <label class="label">
      <span class="label-text">Select Storage Type</span>
    </label>
    <select
      class="select select-bordered w-full max-w-xs"
      bind:value={selectedType}
    >
      {#each typeOptions as storageType}
        <option value={storageType}>
          {storageType}
        </option>
      {/each}
    </select>
  </div>

  <div class="my-3 p-2 border rounded-box flex flex-col gap-2">
    <label class="input-group input-group-md">
      <span class="w-auto">{storageTypes[selectedType]["label"]}</span>
      <input
        type="text"
        placeholder="foo"
        class="input input-md input-bordered flex-grow"
        bind:value={hostDir}
      />
    </label>

    <label class="input-group input-group-md flex">
      <span class="w-auto">In-App Path</span>
      <input
        type="text"
        placeholder="/data"
        class="input input-md input-bordered w-auto flex-grow"
        bind:value={mountDir}
      />
    </label>
  </div>

  {#if error}
    <div class="mb-4">
      <Error {error} action="mounting storage" />
    </div>
  {/if}

  <div class="mt-3">
    <button class="btn" class:loading on:click={dispatchCreateStorage}>
      Submit
    </button>
  </div>
</Modal>
