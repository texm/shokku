<script>
  import {
    commandExecutionIds,
    commandExecutions,
    executionIdDescriptions,
  } from "$lib/stores";
  import Icon from "$common/Icon.svelte";
  import CommandExecution from "./CommandExecution.svelte";

  export let watchingCompleted = false;

  let selectedId;
  let executions = {};
  let ids = [];
  $: if ($commandExecutionIds !== null) {
    let allIds = new Set();
    for (let id in executions) allIds.add(id);
    for (let i in $commandExecutionIds) allIds.add($commandExecutionIds[i]);
    ids = [...allIds];

    let newId;
    for (let i in ids) {
      const id = ids[i];
      if (!executions[id]) newId = id;
      executions[id] = $commandExecutions[id] || [];
    }

    if (newId) selectedId = newId;

    if ($commandExecutionIds.length > 0 && !selectedId) {
      selectedId = $commandExecutionIds[$commandExecutionIds.length - 1];
    }

    watchingCompleted = Object.keys(executions).length > 0;
  }

  const removeSelectedId = () => {
    commandExecutionIds.remove(selectedId);
    let newExec = {};
    let nextId;
    for (const id in executions) {
      if (id !== selectedId) {
        newExec[id] = executions[id];
        nextId = id;
      }
    }
    executions = newExec;
    selectedId = nextId;

    watchingCompleted = Object.keys(executions).length > 0;
  };
</script>

<div
  class="bg-neutral rounded-lg shadow-lg border-info border-2 text-neutral-content w-full h-fit min-h-16 p-2"
>
  <div class="flex flex-row items-center">
    <ul class="menu menu-compact menu-horizontal items-center rounded-box p-1">
      {#each Object.keys(executions) as id}
        <li>
          <a
            class:active={id === selectedId}
            on:click={() => (selectedId = id)}
          >
            {$executionIdDescriptions[id]}
          </a>
        </li>
      {/each}
    </ul>

    <div class="flex-grow" />

    {#if selectedId && executions[selectedId]}
      <div class="items-center">
        <button
          class="btn btn-sm btn-square hover:btn-error"
          on:click={removeSelectedId}
        >
          <Icon type="delete" size="sm" />
        </button>
      </div>
    {/if}
  </div>

  {#if selectedId}
    <CommandExecution status={executions[selectedId]} />
  {/if}
</div>
