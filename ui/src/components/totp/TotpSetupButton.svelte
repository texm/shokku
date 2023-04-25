<script>
  import Modal from "$common/Modal.svelte";
  import { useQuery } from "@sveltestack/svelte-query";
  import { requestGenerateTotp, confirmTotpCode } from "$lib/apis/setup.js";
  import Error from "$common/Error.svelte";
  import Icon from "$common/Icon.svelte";

  export let isSetup;
  export let totpSecret;
  export let recoveryCode;

  let codeConfirmed = false;

  let modalOpen = false;
  const genQuery = useQuery("requestGenerateTotp", requestGenerateTotp, {
    enabled: false,
  });

  const generateTotp = () => {
    genQuery.setEnabled(true);
    modalOpen = true;
  };

  let image;
  $: if ($genQuery.data) {
    image = $genQuery.data["image"];
    totpSecret = $genQuery.data["secret"];
    recoveryCode = $genQuery.data["recovery_code"];
  }

  const copySecret = () => {
    navigator.clipboard.writeText(totpSecret);
  };

  let codeInput;

  let confirmError;
  const tryConfirm = async () => {
    confirmError = null;
    let ok = await confirmTotpCode(totpSecret, codeInput);
    if (ok) {
      codeConfirmed = true;
    } else {
      confirmError = "invalid code";
    }
  };

  const confirmSetup = () => {
    isSetup = true;
    modalOpen = false;
  };
</script>

{#if $genQuery.isError}
  <Error error={$genQuery.error} action="generating totp" />
{/if}
<button
  class="btn btn-primary"
  class:loading={$genQuery.isLoading}
  class:btn-disabled={isSetup}
  class:bg-success={isSetup}
  class:text-success-content={isSetup}
  on:click={generateTotp}
>
  {isSetup ? "TOTP Setup" : "Setup TOTP"}
</button>

<Modal
  name="totp-setup"
  title="TOTP Setup"
  open={modalOpen}
  preventClose={false}
>
  <div class="w-full flex flex-col" class:hidden={codeConfirmed}>
    <div class="">
      <span>Scan:</span>
      <img class="w-[160px] h-[160px]" src="data:image/png;base64,{image}" />

      <span>Or enter secret manually:</span>
      <br />
      <div
        class="btn btn-ghost btn-icon gap-2 rounded-lg p-2"
        on:click={copySecret}
      >
        <Icon type="copy" />
        <code>{totpSecret}</code>
      </div>
    </div>

    <div class="mt-2">
      <span>Then enter code from authenticator:</span>
      <label class="input-group">
        <span class="text-neutral">Code</span>
        <input
          type="text"
          class="input input-bordered"
          on:change={() => (confirmError = null)}
          bind:value={codeInput}
        />
      </label>
    </div>

    <div class="mt-2">
      {#if confirmError}
        <div class="">
          <span class="text-error">{confirmError}</span>
        </div>
      {/if}
      <button
        class="btn btn-primary"
        class:btn-error={!!confirmError}
        disabled={!codeInput}
        on:click={tryConfirm}
      >
        Submit
      </button>
    </div>
  </div>

  <div class="w-full flex flex-col" class:hidden={!codeConfirmed}>
    <span> This is your recovery code, store it somewhere safe. </span>
    <br />
    <div
      class="btn btn-ghost w-fit btn-icon gap-2 rounded-lg p-2"
      on:click={copySecret}
    >
      <Icon type="copy" />
      <code>{recoveryCode}</code>
    </div>
    <button class="btn btn-primary mt-2" on:click={confirmSetup}>
      I have stored this somewhere safe
    </button>
  </div>
</Modal>
