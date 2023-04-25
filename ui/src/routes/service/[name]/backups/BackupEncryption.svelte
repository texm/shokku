<script>
  import { page } from "$app/stores";
  import {
    setServiceBackupEncryption,
    unsetServiceBackupEncryption,
  } from "$lib/api";
  import Error from "$common/Error.svelte";
  import { useMutation } from "@sveltestack/svelte-query";

  export let ready;
  export let enabled;

  const serviceName = $page.params.name;

  let updatingPhrase = false;
  let newPhrase = "";

  const unsetEncryptionMutation = useMutation(
    () => unsetServiceBackupEncryption(serviceName),
    {
      onSuccess: () => {
        enabled = false;
      },
    }
  );

  const handleToggleChange = async () => {
    if (enabled) return;
    $unsetEncryptionMutation.mutate({});
  };

  const updatePhraseMutation = useMutation(
    () => setServiceBackupEncryption(serviceName, newPhrase),
    {
      onSuccess: () => {
        updatingPhrase = false;
      },
    }
  );

  const cancelUpdate = () => {
    updatingPhrase = false;
    newPhrase = "";
  };
</script>

<div class="flex flex-col">
  <div class="">
    <label class="label cursor-pointer w-52">
      <span class="label-text">Encrypt Backups</span>
      <input
        type="checkbox"
        class="toggle"
        disabled={!ready}
        bind:checked={enabled}
        on:change={handleToggleChange}
      />
    </label>
    {#if enabled && !updatingPhrase}
      <button class="btn w-fit" on:click={() => (updatingPhrase = true)}>
        Update Encryption Passphrase
      </button>
    {/if}
  </div>

  {#if updatingPhrase}
    <label class="input-group w-full flex">
      <span class="w-fit">Encryption Passphrase</span>
      <input
        bind:value={newPhrase}
        placeholder=""
        class="input input-bordered flex-grow"
      />
    </label>

    <div class="flex flex-row gap-2 mt-2">
      <button
        class="btn btn-primary w-fit"
        on:click={$updatePhraseMutation.mutate}>Save</button
      >
      <button class="btn w-fit" on:click={cancelUpdate}>Cancel</button>
    </div>
  {/if}
</div>

{#if $updatePhraseMutation.isError}
  <Error
    action="updating encryption passphrase"
    error={$updatePhraseMutation.error}
  />
{/if}

{#if $unsetEncryptionMutation.isError}
  <Error action="unsetting encryption" error={$unsetEncryptionMutation.error} />
{/if}
