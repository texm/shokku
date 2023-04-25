<script>
  import { createEventDispatcher, onMount } from "svelte";
  import { setAppConfig } from "$lib/api";
  import Icon from "$common/Icon.svelte";
  import ConfirmationModal from "$common/ConfirmationModal.svelte";
  import Alert from "$common/Alert.svelte";

  export let vars;
  export let saving = false;
  export let stateDirty = false;
  export let showSaveButton = false;
  export let neutralButtons = false;

  let varsList = [];
  onMount(() => {
    for (let key of Object.keys(vars)) {
      varsList.push([key, vars[key]]);
    }
    varsList = varsList;
  });

  let confirmationModalOpen = false;
  let deletingKey = "";
  let deletingIndex = -1;
  const confirmRemoveKey = (index) => {
    return () => {
      deletingIndex = index;
      deletingKey = varsList[index][0];
      if (varsList[index][0] === "" && varsList[index][1] === "") {
        removeKey();
        return;
      }

      confirmationModalOpen = true;
    };
  };

  const removeKey = () => {
    varsList.splice(deletingIndex, 1);
    varsList = varsList;
    checkStateDirty();

    deletingIndex = -1;
    deletingKey = "";
  };

  const addNewEnvVar = () => {
    varsList.push(["", ""]);
    varsList = varsList;
  };

  const checkStateDirty = () => {
    stateDirty = false;
    if (varsList.length !== Object.keys(vars).length) {
      stateDirty = true;
      return;
    }
    for (let [key, val] of varsList) {
      if (!(key in vars) || vars[key] !== val) {
        stateDirty = true;
        return;
      }
    }
  };

  const dispatch = createEventDispatcher();
  const pairChange = (position) => {
    return (e) => {
      const index = e.target.dataset["index"];
      varsList[index][position] = e.target.value;
      checkStateDirty();
      dispatch("changed", varsList);
    };
  };

  const keyChanged = pairChange(0);
  const valChanged = pairChange(1);
</script>

<div class="w-full grid grid-cols-2 gap-2 mb-2">
  {#each varsList as pair, i}
    <div>
      <label class="input-group input-group-md">
        <span>Key</span>
        <input
          type="text"
          on:change={keyChanged}
          data-index={i}
          value={pair[0]}
          class="input input-md input-bordered w-full"
        />
      </label>
    </div>

    <div class="flex flex-row">
      <div class="flex-grow">
        <label class="input-group input-group-md">
          <span>Value</span>
          <input
            type="text"
            on:change={valChanged}
            data-index={i}
            value={pair[1]}
            class="input input-md input-bordered w-full"
          />
        </label>
      </div>
      <button
        class="btn btn-ghost btn-circle text-error ml-2"
        on:click={confirmRemoveKey(i)}
      >
        <Icon type="remove" />
      </button>
    </div>
  {/each}
</div>

<div class="">
  <button
    class="btn gap-2"
    class:btn-neutral={neutralButtons}
    class:text-neutral={!neutralButtons}
    on:click={addNewEnvVar}
  >
    Add
    <Icon type="add" />
  </button>
  {#if stateDirty && showSaveButton}
    <button
      class="btn btn-primary mt-2 gap-2"
      class:loading={saving}
      on:click={() => dispatch("save", varsList)}
    >
      Save
      <Icon type="save" />
    </button>
  {/if}
</div>

<ConfirmationModal
  name="delete-env-var"
  action="delete '{deletingKey}'"
  bind:open={confirmationModalOpen}
  on:accepted={removeKey}
/>
