<script>
	const bootstrapLink = "https://shokku.dev/install.sh";
	const scriptInstallCommands = [
	  `wget "${bootstrapLink}"`,
		`export SHOKKU_LETSENCRYPT_EMAIL="foo@example.com"`,
	  `sudo bash install.sh`,
	];
</script>

<h2>Requirements</h2>
<h3>Dokku</h3>
<p>
	Shokku requires an underlying Dokku installation.
	You are recommended to manually install & configure it following
	<a href="https://dokku.com/docs/getting-started/installation/">the official documentation</a>.
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

<h3>System Requirements</h3>
<p>
	Please refer to
	<a href="https://dokku.com/docs/getting-started/installation/#system-requirements">
	the operating systems supported by Dokku</a>.
	<br />
	Shokku is lightweight and will consume very little resources.
</p>

<h4>DigitalOcean</h4>
<p>
	For a cheap hosted server to try Shokku out, consider using DigitalOcean via this referral
	link. I get a kickback, and you get $200 in credit for 60 days - why not provision something beefy :)
</p>
<a target="_blank" href="https://www.digitalocean.com/?refcode=2fb615299eca&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge">
	<img src="https://web-platforms.sfo2.digitaloceanspaces.com/WWW/Badge%202.svg" alt="DigitalOcean Referral Badge" />
</a>

<h2>Install Using Script</h2>

<div class="mockup-code">
	{#each scriptInstallCommands as cmd, i}
		<pre data-prefix="{i}" class="m-0 my-1 p-0"><code>{cmd}</code></pre>
	{/each}
</div>

<h2>Manual Installation</h2>
<div class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box">
	<input type="checkbox" />
	<div class="collapse-title text-xl font-medium">
		Instructions
	</div>
	<div class="collapse-content">
		<p>
			If you don't feel comfortable using some random script from the internet,
			the installation can be manually performed fairly easily.
			<br />
			Take a look at the <a href="{bootstrapLink}">bootstrap script</a>, it is
			commented and easy to follow.
		</p>
	</div>
</div>
