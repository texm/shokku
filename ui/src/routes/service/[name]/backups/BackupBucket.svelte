<script>
  import { page } from "$app/stores";
  import { useMutation } from "@sveltestack/svelte-query";

  export let report = {};

  const serviceName = $page.params.name;

  let bucket = "";
  let savedBucket = report["bucket"];

  const updateBucketMutation = useMutation(() => {}, {
    onSuccess: () => (savedBucket = bucket),
  });
</script>

<div class="flex flex-col gap-2 mt-2">
  <label class="input-group w-full flex">
    <span class="w-fit">Backup Bucket</span>
    <input
      bind:value={bucket}
      placeholder="my-s3-bucket"
      class="input input-bordered flex-grow"
      class:input-warning={savedBucket === ""}
    />
  </label>

  {#if bucket !== savedBucket}
    <div class="flex flex-row gap-2">
      <button
        class="btn btn-primary w-fit"
        on:click={$updateBucketMutation.mutate}
      >
        Save
      </button>
      <button class="btn w-fit" on:click={() => (bucket = savedBucket)}>
        Cancel
      </button>
    </div>
  {/if}
</div>
