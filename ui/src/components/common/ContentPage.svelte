<div class="card card-body py-3 px-3 h-full max-h-full text-base-content">
  <div class="hidden sm:flex rounded-lg flex-row w-full max-h-full">
    <slot name="sidebar" />

    <div class="pl-3 h-full w-full flex flex-col overflow-x-auto">
      <div class="w-full">
        <slot name="header" />
      </div>

      <div class="overflow-auto w-full h-full rounded-lg">
        <slot name="content" />
      </div>
    </div>
  </div>

  <!-- mobile nav -->
  <div class="inline-block sm:hidden">
    <slot name="header" />
    <slot name="sidebar" />

    <div class="overflow-auto h-full">
      <div class="rounded-lg w-full h-full overflow-auto pb-6">
        <slot name="content" />
      </div>
    </div>
  </div>
</div>
