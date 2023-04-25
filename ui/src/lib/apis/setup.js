import {createApiClient} from "./client.js";

const setupKeyMiddleware = next => (url, opts) => {
	const key = localStorage.getItem("setup_key");
	if (key) opts.headers = {...opts.headers, "X-Setup-Key": key};
	return next(url, opts);
};

const setupClient = createApiClient("/setup/")
	.middlewares([setupKeyMiddleware]);

export const verifySetupKeyValid = async () => {
	return setupClient.url("verify-key")
		.get()
		.res(r => r.ok)
		.catch(r => false);
}

export const completeCreateAppHandshake = async (code) => {
	return setupClient.url("github/create-app")
		.post({code})
		.json();
}

export const getGithubAppSetupStatus = async () => {
	return setupClient.url("github/status")
		.get()
		.json();
}

export const getGithubAppInstallInfo = async () => {
	return setupClient.url("github/install-info")
		.get()
		.json();
}

export const completeGithubSetup = async (params) => {
	return setupClient.url("github/completed")
		.post(params)
		.res();
}

export const requestGenerateTotp = async () => {
	return setupClient.url("totp/new")
		.post()
		.json()
}

export const confirmTotpCode = async (secret, code) => {
	return setupClient.url("totp/confirm")
		.post({secret, code})
		.json(res => res["valid"])
		.catch(_ => false)
}

export const confirmPasswordSetup = async (opts) => {
	return setupClient.url("password")
		.post(opts)
		.res(res => res.ok)
}