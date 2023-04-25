<script>
  import { getAppProcesses, executeCommandInProcess } from "$lib/api";
  import { page } from "$app/stores";

  import Terminal from "$components/commands/Terminal.svelte";
  import Error from "$common/Error.svelte";

  const appName = $page.params.name;

  let processes;
  let selectedProcess;

  let error;
  let errorAction;
  const initProcessSelection = async () => {
    try {
      let psList = await getAppProcesses(appName);
      processes = psList;
      if (psList.length === 1) {
        selectedProcess = psList[0];
      }
    } catch (e) {
      error = e;
      errorAction = "loading processes";
    }
  };

  let terminalOutput = [];
  const onTerminalInput = async (cmd) => {
    if (cmd === "clear") {
      terminalOutput = [];
      return;
    }
    terminalOutput = [...terminalOutput, { input: cmd }];
    try {
      const res = await executeCommandInProcess(appName, selectedProcess, cmd);
      terminalOutput = [...terminalOutput, res];
    } catch (e) {
      console.error(e);
    }
  };

  initProcessSelection();
</script>

{#if processes && !selectedProcess}
  <div class="form-control w-full max-w-md p-2 flex flex-row gap-2">
    <label class="label">
      <span class="label-text">Process type:</span>
    </label>
    <select
      class="select select-primary select-bordered"
      bind:value={selectedProcess}
    >
      {#each processes as process}
        <option>{process}</option>
      {/each}
    </select>
  </div>
{/if}

{#if selectedProcess}
  <Terminal onInput={onTerminalInput} bind:output={terminalOutput}>
    <div slot="titlebar">
      <span class="text-md"
        >Connected to <span class="font-bold">{selectedProcess}</span></span
      >
    </div>
  </Terminal>
{/if}

{#if error}
  <Error {error} action={errorAction} />
{/if}
