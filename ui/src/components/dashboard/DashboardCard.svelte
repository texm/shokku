<script>
  import Icon from "$common/Icon.svelte";

  export let info;
  export let contentType;

  let state;
  $: if (info) {
    if (info["is_running"]) state = "Running";
    else if (info["is_deployed"]) state = "Deployed";
    else if (!info["is_setup"]) state = "Needs Setup";
  }
</script>

<div class="card shadow-lg hover:shadow-xl bg-neutral">
  <div class="card-body w-full">
    <div class="flex items-center">
      <div class="text-neutral-content text-xl leading-8">
        <div class="flex flex-row gap-3">
          {#if info.type}
            <Icon size="lg" type={info.type} />
          {/if}
          <span>{info.name}</span>
        </div>
      </div>

      <div class="flex-grow" />

      <div>
        {#if contentType === "app" && state}
          <div
            class="flex justify-end"
            class:text-accent={state === "Needs Setup"}
            class:text-warning={state === "Stopped"}
            class:text-success={state === "Running"}
          >
            <div class="inline mr-2">
              <span class="align-middle leading-8">
                {state}
              </span>
            </div>

            <Icon type="circles" size="lg" />
          </div>
        {/if}
      </div>
    </div>

    <div class="mt-4">
      <a class="w-auto" href="/{contentType}/{info.name}">
        <button class="btn gap-2 btn-primary">
          <Icon type="info" />
          <span class="text-lg">View</span>
        </button>
      </a>
    </div>
  </div>
</div>
