<script>
  export let name;
  export let title;
  export let open = false;
  export let preventClose = false;

  let modalId = name + "-modal";

  const closeModal = () => {
    if (preventClose) return;
    open = false;
  };

  const checkCloseModal = (e) => {
    if (e.target.id === modalId) {
      closeModal();
      e.stopPropagation();
    }
  };
</script>

<div
  class="modal modal-middle"
  class:modal-open={open}
  id={modalId}
  on:click={checkCloseModal}
>
  <div class="modal-box p-6 w-auto max-w-full">
    <div class="mb-4">
      <div class="h-full text-base-content">
        <span class="text-xl leading-7">{title}</span>
      </div>
      <button
        class="btn btn-sm btn-circle absolute right-6 top-5"
        class:btn-loading={preventClose}
        for={modalId}
        on:click={closeModal}
      >
        ✕
      </button>
    </div>

    <slot />
  </div>
</div>
