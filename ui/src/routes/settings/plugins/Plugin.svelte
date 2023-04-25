<script>
  export let name;
  export let version;
  export let enabled;
  export let description;
</script>

<div class="card card-compact w-96 shadow-lg bg-base-100">
  <div class="card-body">
    <h2 class="card-title">{name}</h2>
    <span>Version {version}</span>
    <!--span class:text-success={enabled}>{enabled ? "Enabled" : "Disabled"}</span-->
    <p>{description}</p>
  </div>
</div>
