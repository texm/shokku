<script>
  export let lines = [];
  export let prefix = "";
</script>

<div class="mockup-code text-left w-fit">
  {#each lines as line}
    <pre data-prefix={prefix}><code>{line}</code></pre>
  {/each}
</div>
