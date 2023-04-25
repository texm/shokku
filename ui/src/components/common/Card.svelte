<script>
  export let title = "";
  export let actionsRight = false;
</script>

<div class="card shadow-lg bg-base-200">
  <div class="card-body p-4">
    <h2 class="card-title">{title}</h2>
    <slot />
    <div class="card-actions" class:justify-end={actionsRight}>
      <slot name="actions" />
    </div>
  </div>
</div>
