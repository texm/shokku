<script>
  import Modal from "$common/Modal.svelte";
  import { createEventDispatcher } from "svelte";
  import { page } from "$app/stores";

  export let open;
  export let loading;

  const appName = $page.params.name;
  const dispatch = createEventDispatcher();

  let domain;
  const signalAddDomain = () => dispatch("addDomain", { domain });
</script>

<Modal
  name="add-domain"
  title="Add a new domain for {appName}"
  bind:open
  preventClose={loading}
>
  <div class="form-control">
    <label class="input-group">
      <span>Domain</span>
      <input
        type="text"
        placeholder="foo.example.com"
        class="input input-bordered"
        bind:value={domain}
      />
    </label>
  </div>

  <div class="mt-4">
    <button class="btn" class:loading on:click={signalAddDomain}>
      Confirm
    </button>
  </div>
</Modal>
