<script>
  import { completeGithubSetup } from "$lib/api";

  import Error from "$common/Error.svelte";
  import { page } from "$app/stores";
  import { goto, invalidate } from "$app/navigation";
  import { onMount } from "svelte";
  import Card from "$common/Card.svelte";

  let installId = $page.url.searchParams.get("installation_id");
  let error;
  const completeInstall = async () => {
    try {
      const res = await completeGithubSetup({
        code: $page.url.searchParams.get("code"),
        installation_id: Number(installId),
      });
      await invalidate("app:load");
      await goto("/");
    } catch (e) {
      error = e;
    }
  };
  onMount(completeInstall);
</script>

{#if error}
  <Card title="Failed to complete installation">
    <Error action="completing setup" {error} />

    <a
      class="link text-lg my-2"
      href="https://github.com/settings/installations/{installId}"
    >
      Manually uninstall the Github App here
    </a>

    <div slot="actions">
      <a class="link" href="/setup/github">
        <button class="btn btn-ghost">&lt; Restart Installation</button>
      </a>
    </div>
  </Card>
{/if}
