<script>
  import { page } from "$app/stores";
  import {
    setServiceBackupsSchedule,
    unscheduleServiceBackups,
  } from "$lib/api";
  import Error from "$common/Error.svelte";
  import { useMutation } from "@sveltestack/svelte-query";

  export let ready = false;
  export let schedule = "";

  const serviceName = $page.params.name;

  let enableSchedule = !!schedule;
  let newSchedule = "";
  $: if (schedule) newSchedule = schedule;

  const removeScheduleMutation = useMutation(() =>
    unscheduleServiceBackups(serviceName)
  );

  const onToggled = async () => {
    if (enableSchedule) return;
    $removeScheduleMutation.mutate({});
  };

  const updateScheduleMutation = useMutation(
    () => setServiceBackupsSchedule(serviceName, newSchedule),
    {
      onSuccess: () => {
        schedule = newSchedule;
      },
    }
  );
</script>

<div class="flex flex-col">
  <div class="">
    <label class="label cursor-pointer w-52">
      <span class="label-text">Schedule Backups</span>
      <input
        type="checkbox"
        class="toggle"
        disabled={!ready}
        bind:checked={enableSchedule}
        on:change={onToggled}
      />
    </label>
    {#if enableSchedule}
      <label class="input-group w-full flex">
        <span class="w-fit">Crontab Schedule</span>
        <input
          bind:value={newSchedule}
          placeholder="0 3 * * *"
          class="input input-bordered flex-grow"
        />
      </label>

      {#if schedule !== newSchedule}
        <div class="flex flex-row gap-2 mt-2">
          <button
            class="btn btn-primary w-fit"
            on:click={$updateScheduleMutation.mutate}
          >
            Save
          </button>
          <button class="btn w-fit" on:click={() => (newSchedule = schedule)}>
            Cancel
          </button>
        </div>
      {/if}
    {/if}
  </div>
</div>

{#if $removeScheduleMutation.isError}
  <Error action="removing schedule" error={$removeScheduleMutation.error} />
{/if}

{#if $updateScheduleMutation.isError}
  <Error action="updating schedule" error={$updateScheduleMutation.error} />
{/if}
