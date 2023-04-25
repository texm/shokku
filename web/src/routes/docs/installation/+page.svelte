<script>
	const scriptInstallCommands = [
	  `wget "https://shokku.app/install.sh"`,
		`export SHOKKU_LETSENCRYPT_EMAIL="foo@example.com"`,
	  `sudo bash install.sh`,
	]

  const manualInstallCommands = [
	  `todo`,
  ]
</script>

<h2>Prequisites</h2>
<h3>Dokku</h3>
<p>
	Shokku requires an underlying Dokku installation.
	You are recommended to manually install & configure it following
	<a href="https://dokku.com/docs/getting-started/installation/">the official documentation</a>.
	<br />
	If an existing installation is not detected during the setup process, it
	will be automatically installed and configured.
</p>

<h3>SSL</h3>
<p>
	Shokku will only work over HTTPS, so SSL must be configured. This can be via:
</p>

<ul>
	<li>
		<strong>Automatic (LetsEncrypt)</strong> <br />
		If a global email has been set, it will automatically be used. Otherwise,
		ensure the <code>SHOKKU_LETSENCRYPT_EMAIL</code> environment variable is set during
		installation to a valid email.
	</li>
	<li>
		<strong>Manual</strong>
		<br />
		Set the <code>SHOKKU_CERT</code> environment variable to point to a tarball containing the
		certificate contents. Tarball creation instructions can be found
		<a href="https://dokku.com/docs/configuration/ssl/#certificate-setting">here</a>.
	</li>
</ul>

<!--div class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box">
	<input type="checkbox" />
	<div class="collapse-title text-lg font-medium">
		Default installation configuration
	</div>
	<div class="collapse-content">
		<p>nginx, letsencrypt, etc</p>
	</div>
</div-->

<h2>Install Using Script</h2>

<div class="mockup-code">
	{#each scriptInstallCommands as cmd, i}
		<pre data-prefix="{i}" class="m-0 my-1 p-0"><code>{cmd}</code></pre>
	{/each}
</div>

<h2 class="flex items-center gap-2">
	<img alt="digitalocean icon" class="w-8 h-8"
			 src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/digitalocean/digitalocean-original.svg" />
	DigitalOcean 1-Click Install
</h2>
<p>
	TODO: link to digitalocean marketplace
</p>

<h2>Manual Installation</h2>
<div class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box">
	<input type="checkbox" />
	<div class="collapse-title text-xl font-medium">
		Instructions
	</div>
	<div class="collapse-content">
		<p>
			If you don't feel comfortable using some random script from the internet,
			the installation can be manually performed fairly easily:
		</p>
		<div class="mockup-code">
			{#each manualInstallCommands as cmd, i}
				<pre data-prefix="{i}" class="m-0 my-1 p-0"><code>{cmd}</code></pre>
			{/each}
		</div>
	</div>
</div>
