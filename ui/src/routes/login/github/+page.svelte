<script>
  import { onMount } from "svelte";
  import { useQuery } from "@sveltestack/svelte-query";
  import { createStoredState } from "$lib/auth";
  import { getGithubAuthInfo } from "$lib/api";
  import { page } from "$app/stores";

  import Icon from "$common/Icon.svelte";
  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Card from "$common/Card.svelte";

  const host = $page.url.origin;
  const urlQuery = useQuery("getGithubAuthInfo", getGithubAuthInfo);

  let state;
  onMount(() => (state = createStoredState("github_auth_state")));

  let url;
  $: if ($urlQuery.data) {
    let params = new URLSearchParams({
      client_id: $urlQuery.data["client_id"],
      redirect_uri: `${host}/login/github/callback`,
      state: state,
      allow_signup: false,
    });
    url = "https://github.com/login/oauth/authorize?" + params.toString();
  }
</script>

<QueryDataWrapper query={urlQuery} action="github redirect">
  <Card title="Login">
    <a href={url}>
      <button
        class="btn btn-outline w-full gap-2 fill-neutral-content hover:fill-neutral"
      >
        <Icon type="github" />
        Continue with github
      </button>
    </a>
  </Card>
</QueryDataWrapper>
