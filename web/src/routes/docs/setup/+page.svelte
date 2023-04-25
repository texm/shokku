<script>
	const ghAppUrl = "https://docs.github.com/en/apps/creating-github-apps/creating-github-apps/about-apps"
	const manifestUrl = "https://docs.github.com/en/apps/creating-github-apps/creating-github-apps/creating-a-github-app-from-a-manifest";
	// TODO
	const codeManifestUrl = "https://github.com/texm/shokku"
</script>

<h2>Get Setup Key</h2>
<p>
	Once installation is finished, a <i>setup key</i> will be output to
	the console. Copy this, you will need it for the initial web setup.
</p>

<h2>Choose Setup Method</h2>
<h3>Github</h3>
<p>
	Setting up with Github creates a <a href={ghAppUrl}>Github App</a> to
	act on behalf of your installation. This means authentication will use
	Github OAuth, and information about your repositories can be automatically
	fetched via the Github API.
</p>

<p>
	The app is created using a <a href={manifestUrl}>manifest</a> - you can
	view the configuration <a href={codeManifestUrl}>here</a>.

	You can choose to install the app on your account, or for an organisation.
</p>

<h4>Organisation Install</h4>
<p>
	If you install on behalf of an organisation, all members of the organisation
	will be able to authenticate and use Shokku. There is currently no role-based
	access, so ensure you trust your fellow org members as admins!
</p>

<h3>Username/Password</h3>
<p>
	You can also use username/password authentication.
	Currently this is limited to a single user.
</p>

<h4>TOTP 2FA</h4>
<p>
	Use of TOTP is <strong>strongly recommended</strong>.
	A QR code/secret will be generated, along with a recovery code to be used should
	you lose access to your account. If you choose not to enable this during installation,
	it can be configured later in Settings.
</p>