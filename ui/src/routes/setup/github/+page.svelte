<script>
  import { onMount } from "svelte";
  import { createStoredState } from "$lib/auth";
  import { page } from "$app/stores";
  import Card from "$common/Card.svelte";

  const host = $page.url.origin;
  const appUrl =
    host === "http://localhost:5173" ? "https://example.com" : host;
  let manifest = {
    name: `${$page.url.hostname}`,
    url: appUrl,
    hook_attributes: {
      url: `${appUrl}/api/github/events`,
    },
    redirect_url: `${host}/setup/github/created`,
    setup_url: `${host}/setup/github/installed`,
    callback_urls: [
      `${host}/setup/github/installed`,
      `${host}/login/github/callback`,
    ],
    public: false,
    default_permissions: {
      contents: "read",
      members: "read",
    },
    default_events: ["push", "meta"],
  };

  let state;
  onMount(() => (state = createStoredState("github_install_state")));

  let useOrganization = false;
  let organization = "";
  let formUrlBase;

  $: if (useOrganization)
    formUrlBase = `https://github.com/organizations/${organization}`;
  $: if (!useOrganization) formUrlBase = `https://github.com`;
</script>

<Card title="Configure Installation">
  <div class="my-2">
    <div class="form-control w-60">
      <label class="label cursor-pointer">
        <span class="label-text">Install for organization?</span>
        <input type="checkbox" class="toggle" bind:checked={useOrganization} />
      </label>
    </div>

    {#if useOrganization}
      <label class="input-group">
        <span class="text-neutral">Organization Name</span>
        <input
          type="text"
          class="input input-bordered"
          bind:value={organization}
        />
      </label>
    {/if}
  </div>

  <form action="{formUrlBase}/settings/apps/new?state={state}" method="post">
    <input
      class="hidden"
      type="text"
      name="manifest"
      id="manifest"
      value={JSON.stringify(manifest)}
    />
    <input
      class="btn btn-primary"
      disabled={useOrganization && !organization}
      type="submit"
      value="Install"
    />
  </form>
</Card>
