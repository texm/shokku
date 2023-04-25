<script>
  import { updateServiceBackupAuth } from "$lib/api";
  import { page } from "$app/stores";
  import Error from "$common/Error.svelte";
  import { useMutation } from "@sveltestack/svelte-query";
  import Card from "$common/Card.svelte";

  export let setup = false;

  let config = {};

  const serviceName = $page.params.name;

  let managingAuth = !setup;

  let accessKeyId = "";
  let secretKey = "";

  let showOptional = false;
  let region = "";
  let signatureVersion = "";
  let endpointURL = "";

  const resetAuthConfig = () => {
    managingAuth = false;
    accessKeyId = config["access_key_id"];
    secretKey = config["secret_key"];
    region = config["region"];
    signatureVersion = config["signature_version"];
    endpointURL = config["endpoint_url"];
  };

  const onSuccess = (newConfig) => {
    config = newConfig;
    managingAuth = false;
  };
  const updateAuthMutation = useMutation(
    (newConfig) => updateServiceBackupAuth(serviceName, newConfig),
    { onSuccess }
  );

  const saveAuthDetails = async () => {
    const newConfig = {
      access_key_id: accessKeyId,
      secret_key: secretKey,
      region: region,
      signature_version: signatureVersion,
      endpoint_url: endpointURL,
    };
    $updateAuthMutation.mutate(newConfig);
  };
</script>

<div class="w-auto mt-2">
  <div class="flex flex-col gap-2">
    {#if managingAuth}
      <label class="input-group w-full flex">
        <span class="w-fit">Access Key ID</span>
        <input
          bind:value={accessKeyId}
          class="input input-bordered flex-grow"
        />
      </label>

      <label class="input-group w-full flex">
        <span class="w-fit">Secret Access Key</span>
        <input bind:value={secretKey} class="input input-bordered flex-grow" />
      </label>

      <label class="label cursor-pointer w-52">
        <span class="label-text">Show Optional Values</span>
        <input type="checkbox" class="toggle" bind:checked={showOptional} />
      </label>

      {#if showOptional}
        <label class="input-group w-full flex">
          <span class="w-fit">Region</span>
          <input
            bind:value={region}
            placeholder="us-east-1"
            class="input input-bordered flex-grow"
          />
        </label>
        <label class="input-group w-full flex">
          <span class="w-fit">Signature version</span>
          <input
            bind:value={signatureVersion}
            placeholder="s3v4"
            class="input input-bordered flex-grow"
          />
        </label>
        <label class="input-group w-full flex">
          <span class="w-fit">Endpoint URL</span>
          <input
            bind:value={endpointURL}
            placeholder="https://YOURMINIOSERVICE"
            class="input input-bordered flex-grow"
          />
        </label>
      {/if}

      <div class="flex flex-row gap-2">
        <button class="btn btn-primary w-fit" on:click={saveAuthDetails}
          >Save</button
        >
        <button
          class="btn w-fit"
          class:hidden={!setup}
          on:click={resetAuthConfig}>Cancel</button
        >
      </div>
    {:else}
      <button class="btn w-fit" on:click={() => (managingAuth = true)}>
        Configure auth
      </button>
    {/if}
  </div>
</div>

{#if $updateAuthMutation.isError}
  <div class="mt-2">
    <Error action="updating auth details" error={$updateAuthMutation.error} />
  </div>
{/if}
