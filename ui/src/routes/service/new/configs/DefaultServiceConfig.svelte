<script>
  import ServiceConfigOverrideOption from "./config-components/ServiceConfigOverrideOption.svelte";
  import ServiceConfigEnvVarsOption from "./config-components/ServiceConfigEnvVarsOption.svelte";

  export let config;

  const configOptionsCfg = {
    key: "config-options",
    enabled: true,
    labelText: "Extra CLI arguments to pass to the container create command",
    inputText: "args",
  };

  const customEnvCfg = {
    key: "custom-env",
    enabled: true,
    labelText: "Environment variables to start the service with",
  };

  // TODO: dont hardcode enabled to false so clicking back doesnt collapse options
  const optionalConfigs = [
    {
      key: "image",
      enabled: false,
      toggleText: "Use custom image name?",
      inputText: "Image Name",
    },
    {
      key: "image-version",
      enabled: false,
      toggleText: "Use custom image version?",
      inputText: "Image Version",
    },
    {
      key: "memory",
      enabled: false,
      type: "number",
      toggleText: "Limit container memory?",
      inputText: "Memory (MB)",
    },
    {
      key: "password",
      enabled: false,
      toggleText: "Override user-level service password?",
      inputText: "User Password",
    },
    {
      key: "root-password",
      enabled: false,
      toggleText: "Override root-level service password?",
      inputText: "Root Password",
    },
    {
      key: "shm-size",
      enabled: false,
      type: "number",
      toggleText: "Override shared memory for container?",
      inputText: "Memory (MB)",
    },
  ];
</script>

<div class="flex flex-col gap-2">
  <ServiceConfigOverrideOption
    {...configOptionsCfg}
    value={config[configOptionsCfg.key] || ""}
    showToggle={false}
    on:changed
  />
  <ServiceConfigEnvVarsOption
    {...customEnvCfg}
    value={config[customEnvCfg.key] || []}
    on:changed
  />
  {#each optionalConfigs as cfg}
    <ServiceConfigOverrideOption
      {...cfg}
      value={config[cfg.key] || ""}
      on:changed
    />
  {/each}
</div>
