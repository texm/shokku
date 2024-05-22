<script>
  import { fly } from "svelte/transition";
  import { createEventDispatcher } from "svelte";
  import Card from "$common/Card.svelte";

  export let steps = [];
  export let props = {};
  export let data = {};
  export let loading;
  export let confirmButtonText = "Confirm";

  let finishedStep = {};
  let currentStep = 0;

  const canSetStep = (step) => {
    if (loading) return false;
    if (step <= 1) return finishedStep[0];
    return finishedStep[step - 1] && canSetStep(step - 1);
  };

  const updateStepStatus = (step, { complete }) =>
    (finishedStep[step] = complete);

  const maybeSetStep = (step) => {
    if (!canSetStep(step)) return;
    currentStep = step;
  };

  const dispatch = createEventDispatcher();
  const confirmClicked = () => dispatch("complete");
</script>

<div class="flex flex-row gap-4 max-h-full w-full">
  <div class="flex bg-base-200 shadow-xl rounded-lg py-4 px-2 h-fit">
    <ul class="steps steps-vertical">
      {#each steps as step, i}
        <li
          class="step px-4"
          class:cursor-pointer={!loading && finishedStep[i]}
          class:step-primary={finishedStep[i]}
          on:click={() => maybeSetStep(i)}
          data-content={i === currentStep ? "●" : null}
        >
          <span>{step.label}</span>
        </li>
      {/each}
    </ul>
  </div>

  <div
    class="flex flex-col bg-base-200 p-4 rounded-lg shadow-xl flex-grow max-h-full"
  >
    <div class="overflow-auto" in:fly={{ y: 100, duration: 250 }}>
      <svelte:component
        this={steps[currentStep].component}
        {props}
        on:statusChange={(e) => updateStepStatus(currentStep, e.detail)}
        bind:data
      />
    </div>

    <div class="h-fit">
      <slot name="errors" />

      <div class="flex gap-2 mt-2 h-fit">
        <button
          class="btn"
          class:btn-disabled={loading || currentStep === 0}
          on:click={() => maybeSetStep(currentStep - 1)}
        >
          Previous
        </button>
        <button
          class="btn btn-primary"
          class:hidden={currentStep + 1 === steps.length}
          class:btn-disabled={!finishedStep[currentStep]}
          on:click={() => maybeSetStep(currentStep + 1)}
        >
          Next
        </button>
        <button
          class="btn btn-primary"
          class:hidden={currentStep + 1 !== steps.length}
          class:loading
          on:click={() => confirmClicked()}
        >
          {confirmButtonText}
        </button>
      </div>
    </div>
  </div>
</div>
