<script>
  import { page } from "$app/stores";
  import { getServiceBackupReport } from "$lib/api";
  import { useQuery } from "@sveltestack/svelte-query";

  import QueryDataWrapper from "$common/QueryDataWrapper.svelte";
  import Cards from "$common/Cards.svelte";
  import Card from "$common/Card.svelte";

  import BackupButton from "./BackupButton.svelte";
  import BackupBucket from "./BackupBucket.svelte";
  import BackupAuth from "./BackupAuth.svelte";
  import BackupSchedule from "./BackupSchedule.svelte";
  import BackupEncryption from "./BackupEncryption.svelte";
  import Alert from "$common/Alert.svelte";

  export let data;

  const serviceName = $page.params.name;
  const queryKey = [{ serviceName }, "getServiceBackupReport"];
  const serviceBackupReport = useQuery(queryKey, () =>
    getServiceBackupReport(serviceName)
  );

  let report = {};
  let ready;
  $: if ($serviceBackupReport.isSuccess) {
    report = $serviceBackupReport.data;
    ready = report["auth_set"] && report["bucket"] !== "";
  }
</script>

<QueryDataWrapper query={serviceBackupReport} action="loading backup info">
  <Cards>
    <Card title="Run Backup">
      {#if !ready}
        <Alert
          type="warning"
          message="Configure bucket & authentication first"
        />
      {/if}
      <BackupButton {ready} />
    </Card>

    <Card title="Bucket">
      <BackupBucket {report} />
    </Card>

    <Card title="Authentication">
      <BackupAuth setup={report["auth_set"]} />
    </Card>

    <Card title="Schedule">
      <BackupSchedule schedule={report["schedule"]} {ready} />
    </Card>

    <Card title="Encryption">
      <BackupEncryption enabled={report["encryption_set"]} {ready} />
    </Card>
  </Cards>
</QueryDataWrapper>
