<script>
  import { useMutation } from "@sveltestack/svelte-query";
  import { createService } from "$lib/api";
  import { goto } from "$app/navigation";

  import Steps from "$common/Steps.svelte";
  import Error from "$common/Error.svelte";

  import SelectService from "./steps/SelectService.svelte";
  import ConfigureService from "./steps/ConfigureService.svelte";
  import Confirm from "./steps/Confirm.svelte";

  const steps = [
    { label: "Select a service type", component: SelectService },
    { label: "Configure service", component: ConfigureService },
    { label: "Confirm", component: Confirm },
  ];

  const createServiceMutation = useMutation(
    ({ name, type, config }) => createService(name, type, config),
    { onSuccess: (_, { name }) => goto(`/service/${name}`) }
  );

  let props = {};
  let data = {};

  const tryCreateService = () => {
    let type = data.selectedService.type;
    let name = data.config[type].name;
    let cfg = Object.keys(data.config[type])
      .filter((k) => k !== "name")
      .reduce((cur, key) => ({ ...cur, [key]: data.config[type][key] }), {});
    $createServiceMutation.mutate({ name: name, type: type, config: cfg });
  };
</script>

<div class="p-4 max-h-full h-full">
  <Steps
    {steps}
    {props}
    {data}
    loading={$createServiceMutation.isLoading}
    on:complete={tryCreateService}
  />

  {#if $createServiceMutation.isError}
    <Error action="creating service" error={$createServiceMutation.error} />
  {/if}
</div>
