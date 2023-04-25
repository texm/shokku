<script>
  import Icon from "$common/Icon.svelte";
  import { createEventDispatcher } from "svelte";

  export let isAuthenticated = false;
  export let showExecutionButton = false;
  export let executionWindowVisible = false;

  const authenticatedLinks = [
    { href: "/settings", text: "Settings" },
    { href: "/logout", text: "Logout" },
  ];
  const linkStyle = "px-3 py-2 btn btn-ghost";

  const dispatch = createEventDispatcher();
  const executionButtonClicked = () => dispatch("executionButtonClicked");
</script>

<nav
  class="bg-neutral text-neutral-content flex flex-col md:flex-row items-center w-full justify-between p-4 mx-auto"
>
  <a
    class="hover:font-bold inline-flex items-center justify-center btn btn-ghost normal-case text-xl"
    href="/">shokku</a
  >

  {#if showExecutionButton}
    <button
            class="btn gap-2 hover:btn-secondary"
            class:btn-active={executionWindowVisible}
            on:click={executionButtonClicked}
    >
      <Icon type="file-text" />
      command output
    </button>
  {/if}

  <div class="flex flex-col">
    <ul class="flex flex-col md:flex-row items-center space-x-2 text-sm font-medium">
      {#if isAuthenticated}
        {#each authenticatedLinks as link}
          <li>
            <a class={linkStyle} href={link["href"]}>{link["text"]}</a>
          </li>
        {/each}
      {/if}
    </ul>
  </div>
</nav>
