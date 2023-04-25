<script>
  import { page } from "$app/stores";
  import { goto, invalidate } from "$app/navigation";
  import { completeGithubAuth } from "$lib/api";
  import { checkStoredState } from "$lib/auth";

  import ErrorAlert from "$common/Error.svelte";
  import Loader from "$common/Loader.svelte";
  import Card from "$common/Card.svelte";

  let loading = true;
  let error;
  const completeAuth = async () => {
    const url = $page.url;
    const state = url.searchParams.get("state");
    const code = url.searchParams.get("code");
    if (!checkStoredState("github_auth_state", state)) {
      error = new Error("state does not match");
      return;
    }
    try {
      let redirectUrl = url.origin + url.pathname;
      await goto($page.url.pathname, { replaceState: true });
      await completeGithubAuth(code, redirectUrl);
      await invalidate("app:load");
      await goto("/");
    } catch (e) {
      console.log("error completing login", e);
      error = e;
    } finally {
      loading = false;
    }
  };
  completeAuth();
</script>

<Card title="Completing Authentication">
  {#if loading}
    <div class="flex flex-row gap-2 items-center">
      <Loader />
      <span class="text-lg">Validating</span>
    </div>
  {:else if error}
    <ErrorAlert {error} action="validating credentials" />

    <div class="mt-2">
      <a class="link m-2" href="/auth/github">
        <button class="btn btn-ghost">&lt; Back to Login</button>
      </a>
    </div>
  {/if}
</Card>
