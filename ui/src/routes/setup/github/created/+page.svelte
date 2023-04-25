<script>
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";
  import { completeCreateAppHandshake } from "$lib/api";
  import { checkStoredState } from "$lib/auth";

  import ErrorAlert from "$common/Error.svelte";
  import Card from "$common/Card.svelte";
  import Error from "$common/Error.svelte";

  let error;
  const verifyState = async () => {
    const state = $page.url.searchParams.get("state");
    const code = $page.url.searchParams.get("code");
    if (!checkStoredState("github_install_state", state)) {
      error = new Error("state does not match");
      return;
    }
    try {
      await completeCreateAppHandshake(code);
      await goto("/setup/github/install");
    } catch (e) {
      error = e;
    }
  };
  onMount(verifyState);
</script>

{#if error}
  <Card title="Failed to create app">
    <Error action="completing app creation" {error} />

    <div slot="actions">
      <a class="link" href="/setup/github">
        <button class="btn btn-ghost">&lt; Restart Installation</button>
      </a>
    </div>
  </Card>
{/if}
