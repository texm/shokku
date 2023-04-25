<script>
  import { createEventDispatcher } from "svelte";
  import Icon from "./Icon.svelte";

  const dispatch = createEventDispatcher();

  export let message;
  export let type = "error";
  export let canDismiss = false;

  const clicked = () => {
    if (canDismiss) message = "";
  };
</script>

<div class="mt-2" class:hidden={!message || message === ""}>
  <div
    class="alert alert-error shadow-lg"
    class:cursor-pointer={canDismiss}
    on:click={clicked}
  >
    <div>
      <Icon {type} />
      <span class="text-error-content">{message}</span>
    </div>
  </div>
</div>
