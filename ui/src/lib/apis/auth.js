import {createApiClient} from "./client.js";

export const authClient = createApiClient("/auth/");

export const doLogin = async (authDetails) => {
	return await authClient.url("login")
		.post(authDetails)
		.json()
		.catch(_ => false)
}

export const refreshAuthToken = async () => {
	return authClient.url("refresh").post().res();
}

export const doLogout = async () => {
	return await authClient.url("logout")
		.post()
		.res(r => r.ok);
}

export const fetchAuthDetails = async () => {
	return await authClient.url("details")
		.get()
		.forbidden(err => {})
		.json(r => r)
		.catch(err => {})
}

export const getGithubAuthInfo = async () => {
	return await authClient.url("github")
		.get()
		.json();
}

export const completeGithubAuth = async (code, redirectUrl) => {
	return await authClient.url("github/auth")
		.post({code, "redirect_url": redirectUrl})
		.res(r => r.ok)
}
