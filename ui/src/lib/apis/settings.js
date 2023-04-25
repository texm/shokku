import {createApiClient} from "./client.js";

const settingsClient = createApiClient("/settings/");

export const getSettings = async () => {
	return await settingsClient.url("basic")
		.get()
		.json(res => res["settings"]);
}

export const getLetsEncryptStatus = async () => {
	return await settingsClient.url("letsencrypt")
		.get()
		.json(res => res);
}

export const getUsers = async () => {
	return await settingsClient.url("users")
		.get()
		.json(res => res["users"]);
}

export const getSSHKeys = async () => {
	return await settingsClient.url("ssh-keys")
		.get()
		.json(res => res["keys"]);
}

export const getGlobalDomainsList = async () => {
	return await settingsClient.url("domains")
		.get()
		.json(res => res["domains"]);
}

export const addGlobalDomain = async (domain) => {
	return await settingsClient.url("domains")
		.post({domain})
		.res(res => res.ok);
}

export const removeGlobalDomain = async (domain) => {
	return await settingsClient.url("domains")
		.query({domain})
		.delete()
		.res(res => res.ok);
}

export const getVersions = async () => {
	return await settingsClient.url("versions")
		.get()
		.json(res => res);
}

export const getNetworksList = async (appName) => {
	return await settingsClient.url("networks")
		.get()
		.json(res => res["networks"]);
}

export const getPlugins = async (appName) => {
	return await settingsClient.url("plugins")
		.get()
		.json(res => res["plugins"]);
}

export const getDockerRegistryReport = async () => {
    return await settingsClient.url("registry")
        .get()
        .json(res => res);
}

export const setDockerRegistry = async ({server, username, password}) => {
    return await settingsClient.url("registry")
        .post({server, username, password})
        .res(res => res.ok);
}

export const addGitAuth = async ({host, username, password}) => {
    return await settingsClient.url("git-auth")
        .post({host, username, password})
        .res(res => res.ok);
}

export const removeGitAuth = async (host) => {
    return await settingsClient.url("git-auth")
        .post({host})
        .res(res => res.ok);
}