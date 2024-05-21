<script>
  import { afterUpdate, beforeUpdate } from "svelte";

  export let onInput;
  export let output = [];

  let bg = "bg-base-200";
  let input = "";

  const defaultAppend = (lines) => {
    output = [...output, ...lines];
  };

  const echoInput = (s) => {
    defaultAppend({ output: s });
  };

  $: if (!onInput) onInput = echoInput;

  const KEY_ENTER = 13;
  const inputKeyPress = (e) => {
    if (e.charCode !== KEY_ENTER) return;

    onInput(input);
    input = "";
  };

  let autoscroll;
  let scrollDiv;
  beforeUpdate(() => {
    autoscroll =
      scrollDiv &&
      scrollDiv.offsetHeight + scrollDiv.scrollTop >
        scrollDiv.scrollHeight - 20;
  });

  afterUpdate(() => {
    if (autoscroll) scrollDiv.scrollTo(0, scrollDiv.scrollHeight);
  });
</script>

<div class="flex flex-col">
  <div
    class="flex flex-shrink items-center px-2 w-full h-8 bg-neutral-focus text-neutral-content"
  >
    <slot name="titlebar" />
  </div>

  <div
    class="flex-grow max-w-80 h-80  bg-base-content p-2"
    bind:this={scrollDiv}
  >
    <div class="whitespace-pre-wrap flex flex-col">
      {#each output as line, i}
        {#if line.input}
          <span class="text-base-100/70 mt-2" data-prefix=">"
            ><code class="text-inherit">{line.input}</code></span
          >
        {/if}
        {#if line.output}
          <span class="text-neutral-content"
            ><code class="text-inherit">{line.output}</code></span
          >
        {/if}
        {#if line.error}
          <span class="text-warning"
            ><code class="text-inherit">{line.error}</code></span
          >
        {/if}
      {/each}
    </div>
  </div>

  <div
    class="flex-shrink px-2 flex flex-row gap-2 min-h-12 w-full h-12 bg-neutral-focus text-neutral-content"
  >
    <span class="text-md self-center">></span>
    <input
      type="text"
      on:keypress={inputKeyPress}
      bind:value={input}
      class="input px-0 bg-neutral-focus rounded-t-none py-1 input-bordered w-full"
    />
  </div>
</div>

<style lang="postcss">
  span[data-prefix]:before {
    content: attr(data-prefix);
    width: 2rem;
    margin-right: 1ch;
    opacity: 0.5;
  }
</style>
