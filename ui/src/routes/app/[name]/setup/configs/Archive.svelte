<script>
  import { createEventDispatcher } from "svelte";
  import Dropzone from "svelte-file-dropzone";
  import Icon from "$common/Icon.svelte";

  export let options;

  if (!options) options = { archive: null };

  const dispatch = createEventDispatcher();
  const checkOptionsValid = () => {
    dispatch("validityChange", { valid: !!options["archive"] });
  };

  function handleFilesSelect(e) {
    const { acceptedFiles } = e.detail;
    if (acceptedFiles.length > 0) {
      const file = acceptedFiles[0];
      if (file.type === "application/x-gzip" && !file.name.endsWith(".tar.gz"))
        return;
      archiveChosen(file);
    }
  }

  const archiveChosen = (archiveFile) => {
    options["archive"] = archiveFile;
    checkOptionsValid();
  };

  let hovering = false;
  const dragEnter = () => (hovering = true);
  const dragLeave = () => (hovering = false);
</script>

<div
  class="p-6 items-center rounded-lg cursor-pointer text-neutral-content bg-neutral hover:bg-neutral-focus"
  class:bg-neutral-focus={hovering}
>
  <Dropzone
    disableDefaultStyles="true"
    accept=".zip, application/gzip, .gz"
    on:drop={handleFilesSelect}
    on:dragenter={dragEnter}
    on:dragleave={dragLeave}
  >
    <div class="flex flex-col items-center p-6">
      <Icon type="upload" size="lg" />
      <p class="my-1">
        <span class="font-semibold">Click to upload</span> or drag and drop
      </p>
      <p class="text-sm">.zip or .tar.gz</p>
    </div>
  </Dropzone>
</div>

{#if options["archive"]}
  <div class="mt-2">
    <span
      >Uploading <span class="font-bold">{options["archive"].name}</span></span
    >
  </div>
{/if}
