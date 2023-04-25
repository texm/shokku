<script>
  import { goto } from "$app/navigation";
  import { doLogin } from "$lib/api";

  import Error from "$common/Error.svelte";
  import Card from "$common/Card.svelte";

  let username = "";
  let password = "";
  let totp = "";

  let showTotp = false;

  let loading;
  let error;
  const tryLogin = async () => {
    loading = true;
    error = null;
    let res = await doLogin({ username, password, totp });
    loading = false;

    if (!res || !res.success) {
      error = "failed to login";
      return;
    }
    if (res["needs_totp"]) {
      showTotp = true;
      return;
    }

    await goto("/");
  };
</script>

<Card title={showTotp ? "Login - 2FA" : "Login"}>
  <div class:hidden={showTotp}>
    <div class="form-control">
      <label class="label" for="username">
        <span class="label-text">Username</span>
      </label>
      <input
        bind:value={username}
        id="username"
        type="text"
        placeholder="username"
        class="input input-bordered"
      />
    </div>

    <div class="form-control">
      <label class="label" for="password">
        <span class="label-text">Password</span>
      </label>
      <input
        bind:value={password}
        id="password"
        type="password"
        placeholder="password"
        class="input input-bordered"
      />
    </div>
  </div>

  <div class:hidden={!showTotp}>
    <div class="form-control">
      <label class="label" for="totp">
        <span class="label-text">TOTP Code</span>
      </label>
      <input
        bind:value={totp}
        id="totp"
        type="text"
        class="input input-bordered"
      />
    </div>
  </div>

  <div slot="actions" class="form-control mt-2 gap-2">
    {#if error}
      <Error action="logging in" {error} />
    {/if}

    <button class="btn btn-primary" class:loading on:click={tryLogin}>
      Submit
    </button>
  </div>
</Card>
