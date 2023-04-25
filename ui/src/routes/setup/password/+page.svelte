<script>
  import { goto, invalidate } from "$app/navigation";
  import { confirmPasswordSetup } from "$lib/api";
  import Card from "$common/Card.svelte";
  import TotpSetupButton from "$components/totp/TotpSetupButton.svelte";
  import { useMutation } from "@sveltestack/svelte-query";

  let username = "admin";
  let password = "";
  let enableTotp = true;
  let totpSecret;
  let recoveryCode;

  let totpSetup = false;

  const onSuccess = async () => {
    await invalidate("app:load");
    await goto("/");
  };

  const setupMut = useMutation(confirmPasswordSetup, { onSuccess });

  const submitClicked = () => {
    const opts = {
      username: username,
      password: password,
      enable_2fa: enableTotp,
      totp_secret: totpSecret,
      recovery_code: recoveryCode,
    };
    $setupMut.mutate(opts);
  };

  let detailsValid;
  $: detailsValid =
    username && password.length >= 8 && (totpSetup || !enableTotp);
</script>

<Card title="Create a user">
  <label class="input-group">
    <span>Username</span>
    <input
      type="text"
      class="input input-bordered"
      class:input-warning={!username}
      bind:value={username}
    />
  </label>

  <span>Must be >8 characters</span>

  <label class="input-group">
    <span>Password</span>
    <input
      type="password"
      class="input input-bordered"
      class:input-warning={password.length < 8}
      bind:value={password}
    />
  </label>

  <div class="form-control">
    <label class="label w-60 cursor-pointer">
      <span class="label-text">Enable 2FA</span>
      <input
        type="checkbox"
        class="toggle"
        disabled={totpSetup}
        bind:checked={enableTotp}
      />
    </label>

    <div class:hidden={!enableTotp}>
      <TotpSetupButton
        bind:isSetup={totpSetup}
        bind:totpSecret
        bind:recoveryCode
      />
    </div>
  </div>

  <div class="actions">
    <button
      class="btn btn-primary"
      disabled={!detailsValid}
      on:click={submitClicked}
    >
      Submit
    </button>
    <button class="btn btn-ghost mt-2" on:click={() => goto("/setup")}>
      Cancel
    </button>
  </div>
</Card>
