<script>
	import {base} from "$app/paths";

	const gitPushUrl = "https://dokku.com/docs/deployment/application-deployment/#deploy-the-app"
	const remoteRepoDocs = "https://dokku.com/docs/deployment/methods/git/#initializing-an-app-repository-from-a-remote-repository"
	const dockerImageDocs = "https://dokku.com/docs/deployment/methods/git/#initializing-an-app-repository-from-a-docker-image"
</script>

<h2>Creating an App</h2>
<p>
	Once setup, the initial page you see is the <i>Dashboard</i>. From
	here you can create apps and services, or view any existing ones.
	After you have chosen a valid name for your app, you will proceed to
	setup. Currently apps can be setup in three different ways:
</p>

<h2>Create New Git Repo</h2>
<p>
	This option creates a git repo on the host, and sets a branch to deploy
	from. Once configured, you will need to add the remote to your local git
	repo, and push to deploy following <a href={gitPushUrl}>these instructions</a>.
</p>

<h2>Sync Existing Git Repo</h2>
<p>
	This option configures a remote repository to clone using a specific
	<i>git ref</i> (eg the <i>main</i> branch, or a specific commit), which can
	then be automatically built and deployed. See the Dokku docs on this method
	<a href={remoteRepoDocs}>here</a>.
	<br />
	If the repo is private, ensure you have configured access in the Settings page.
</p>

<h2>Docker Image</h2>
<p>
	This option enables pulling a Docker image. See the Dokku docs for caveats
	<a href={dockerImageDocs}>here</a>.
	<br />
	If the image is held in a private registry, ensure you have configured access in the
	Settings page.
</p>