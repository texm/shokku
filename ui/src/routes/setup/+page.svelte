<script>
  import { verifySetupKeyValid } from "$lib/api";
  import { setSetupKey } from "$lib/auth";
  import { useQuery, useQueryClient } from "@sveltestack/svelte-query";

  import Icon from "$common/Icon.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Card from "$common/Card.svelte";
  import Alert from "$common/Alert.svelte";

  let setupKey = "";

  const queryClient = useQueryClient();
  const keyValidQuery = useQuery(["verifySetupKeyValid"], verifySetupKeyValid);

  let isValid;
  let attemptFailed;
  $: if ($keyValidQuery.data) isValid = $keyValidQuery.data === true;

  const verifyClicked = async () => {
    if (!setupKey) return;
    setSetupKey(setupKey);
    await queryClient.invalidateQueries(["verifySetupKeyValid"]);
    attemptFailed = !isValid;
  };
</script>

<QueryDataWrapper query={keyValidQuery}>
  <Card title="Setup">
    <div class="p-2" class:hidden={isValid}>
      <span class="text-lg">Enter key to begin</span>
      <label class="input-group mt-2">
        <span>Setup Key</span>
        <input type="text" class="input input-bordered" bind:value={setupKey} />
      </label>

      {#if attemptFailed}
        <Alert type="error" message="Invalid Key" />
      {/if}

      <button class="btn btn-primary mt-2" on:click={verifyClicked}>
        Verify
      </button>
    </div>

    <div class="flex flex-col gap-2" class:hidden={!isValid}>
      <a href="/setup/github">
        <button
          class="btn w-full gap-2 fill-neutral-content hover:fill-primary-content"
        >
          <Icon type="github" />
          use github authentication
        </button>
      </a>

      <a href="/setup/password">
        <button class="btn btn-outline w-full gap-2 hover:fill-primary">
          <Icon type="key" />
          use password authentication
        </button>
      </a>
    </div>
  </Card>
</QueryDataWrapper>
