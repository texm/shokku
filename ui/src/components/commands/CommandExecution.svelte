<script>
  import Loader from "$common/Loader.svelte";

  export let status = {};

  let msgContainer;
  let finished = false;
  let output = [];
  let bg = "bg-base-200";
  $: if (status) {
    let outputPrev = output.length;
    output = status["output"] || [];
    finished = status["finished"];
    if (status["finished"] && !status["success"]) {
      bg = "bg-error";
    }

    if (msgContainer) {
      msgContainer.scroll({
        top: msgContainer.scrollHeight,
      });
    }
  }
</script>

<div class="max-w-80 max-h-96 overflow-x-scroll" bind:this={msgContainer}>
  {#each output as line, i}
    <pre data-prefix={i} class:text-warning={line["type"] === "stderr"}><code
        >{line["message"]}</code
      ></pre>
  {/each}

  {#if !finished}
    <div class="m-4">
      <Loader />
    </div>
  {/if}
</div>

<style lang="postcss">
  pre[data-prefix]:before {
    content: attr(data-prefix);
    width: 2rem;
    margin-right: 2ch;
    opacity: 0.5;
  }
</style>
