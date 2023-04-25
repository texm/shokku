<script>
  import Icon from "$components/common/Icon.svelte";

  export let action;
  export let error;
  export let errorMessage = null;

  let displayMessage = errorMessage;
  $: if (error && !errorMessage) {
    try {
      let parsed = JSON.parse(error.message);
      if (parsed.message) {
        let code = parsed.error.split(",")[0].substring(5, 8);
        displayMessage = `HTTP ${code}: ${parsed.message}`;
      } else {
        displayMessage = `${parsed.type} error`;
      }
    } catch {
      displayMessage = error;
    }
  }
</script>

<div class="card bg-error shadow-xl my-2">
  <div class="card-body text-error-content">
    <div class="grid grid-cols-3 items-center mb-4">
      <div class="col-span-2 text-neutral-content">
        <span class="text-xl leading-8">Error {action}</span>
      </div>

      <div class="flex justify-end">
        <Icon type="warning" />
      </div>
    </div>
    <p>{displayMessage}</p>
  </div>
</div>
