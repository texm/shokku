<script>
  import Modal from "./Modal.svelte";
  import { createEventDispatcher } from "svelte";

  export let name;
  export let title = "Confirm";
  export let action = "foo the bar";
  export let open = false;
  export let doingAction = false;
  export let extraOption = "";

  let extraOptionChecked = false;

  const dispatch = createEventDispatcher();

  const close = () => (open = false);
  const accept = () => {
    dispatch("accepted", {
      extraOptionChecked,
    });
    close();
  };
</script>

<Modal {name} {title} bind:open preventClose={doingAction}>
  <div class="mb-4">
    <span class="text-md">Are you sure you want to {action}?</span>

    {#if extraOption}
      <div
        class="form-control w-52 border-2 border-base-200 rounded-lg mt-2 p-2"
      >
        <label class="label cursor-pointer">
          <span class="label-text">{extraOption}</span>
          <input
            type="checkbox"
            bind:checked={extraOptionChecked}
            class="checkbox"
          />
        </label>
      </div>
    {/if}
  </div>

  <button class="btn btn-primary" class:loading={doingAction} on:click={accept}
    >Yes</button
  >
  <button
    class="btn btn-secondary btn-ghost"
    class:loading={doingAction}
    on:click={close}>No</button
  >
</Modal>
