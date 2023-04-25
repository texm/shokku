<script>
  export let logs;
</script>

<div class="overflow-x-scroll h-full bg-neutral px-4 rounded-lg">
  <div class="py-4 h-full text-neutral-content">
    {#each logs as line, i}
      <pre data-prefix={i}><code>{line.trim()}</code></pre>
    {/each}
  </div>
</div>

<style lang="postcss">
  pre[data-prefix]:before {
    content: attr(data-prefix);
    width: 2rem;
    margin-right: 2ch;
    opacity: 0.5;
  }
</style>
