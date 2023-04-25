import {createApiClient} from "./client.js";

const appsClient = createApiClient("/apps/");

export const getAppsList = async () => {
	return await appsClient.url("list")
		.get()
		.json(res => res["apps"]);
}

export const getAllAppsOverview = async () => {
    return await appsClient.url("report")
        .get()
        .json(res => res["apps"]);
}

export const getAppOverview = async (appName) => {
	return await appsClient.url("overview")
		.query({"name": appName})
		.get()
		.json();
}

export const getAppIsSetup = async (appName) => {
	return await appsClient.url("setup")
		.query({"name": appName})
		.get()
		.json(res => res["is_setup"]);
}

export const createApp = async (appName) => {
	return appsClient.url("create")
		.post({"name": appName})
		.res(res => res.ok);
}

export const destroyApp = async (appName) => {
	return await appsClient.url("destroy")
		.post({"name": appName})
		.res(res => res.ok);
}

export const startApp = async (appName) => {
	return appsClient.url("start")
		.post({"name": appName})
		.json(res => res["execution_id"]);
}

export const stopApp = async (appName) => {
	return appsClient.url("stop")
		.post({"name": appName})
		.json(res => res["execution_id"]);
}

export const restartApp = async (appName) => {
	return appsClient.url("restart")
		.post({"name": appName})
		.json(res => res["execution_id"]);
}

export const rebuildApp = async (appName) => {
	return appsClient.url("rebuild")
		.post({"name": appName})
		.json(res => res["execution_id"]);
}

export const getAppInfo = async (appName) => {
	return await appsClient.url("info")
		.query({"name": appName})
		.get()
		.json(res => res["info"]);
}

export const renameApp = async (curName, newName) => {
	return await appsClient.url("rename")
		.post({"current_name": curName, "new_name": newName})
		.res(res => res.ok);
}

export const setupAppAsync = async (appName, source, options) => {
	const body = {"name": appName, ...options}
	let req = appsClient.url("setup/" + source)

	if (source === "upload-archive")
		return await req.formData(body).post().json(res => res["execution_id"]);

	return await req.post(body).json(res => res["execution_id"]);
}

export const setupApp = async (appName, source, options) => {
	return await appsClient.url("setup/" + source)
		.post({"name": appName, ...options})
		.res(res => res.ok);
}

export const getAppSetupConfig = async (appName) => {
	return await appsClient.url("setup/config")
		.query({"name": appName})
		.get()
		.json();
}

export const listAppServices = async (appName) => {
	return await appsClient.url("services")
		.query({name: appName})
		.get()
		.json(res => res["services"]);
}

export const getAppDeployChecksReport = async (appName) => {
	return await appsClient.url("deploy-checks")
		.query({"name": appName})
		.get()
		.json(res => res);
}

export const setAppDeployChecksState = async (appName, state) => {
	return await appsClient.url("deploy-checks")
		.post({
			"name": appName,
			"state": state,
		})
		.res(res => res.ok);
}

export const setAppProcessDeployChecksState = async (appName, process, state) => {
	return await appsClient.url("process/deploy-checks")
		.post({
			"name": appName,
			"process": process,
			"state": state,
		})
		.res(res => res.ok);
}

export const getAppProcesses = async (appName) => {
	return await appsClient.url("process/list")
		.query({"name": appName})
		.get()
		.json(res => res["processes"]);
}

export const getAppProcessReport = async (appName) => {
	return await appsClient.url("process/report")
		.query({"name": appName})
		.get()
		.json();
}

export const setAppProcessResources = async (appName, process, resources) => {
	return await appsClient.url("process/resources")
		.post({"name": appName, "process": process, ...resources})
		.res(res => res.ok);
}

export const getAppProcessScale = async (appName) => {
	return await appsClient.url("process/scale")
		.query({"name": appName})
		.get()
		.json(res => res["scale"]);
}

export const setAppProcessScale = async (appName, processName, scale) => {
	return await appsClient.url("process/scale")
		.post({"name": appName, "process": processName, "scale": scale})
		.json(res => res["execution_id"]);
}

export const getAppNetworksReport = async (appName) => {
	return await appsClient.url("networks")
		.query({"name": appName})
		.get()
		.json(res => res);
}

export const updateAppNetworks = async (appName, config) => {
	return await appsClient.url("networks")
		.post({"name": appName, ...config})
		.res(res => res.ok);
}

export const getAppDomainsReport = async (appName) => {
	return await appsClient.url("domains")
		.query({"name": appName})
		.get()
		.json(res => res);
}

export const setAppDomainsEnabled = async (appName, enabled) => {
	return await appsClient.url("domains/state")
		.post({"name": appName, "enabled": enabled})
		.res(res => res.ok);
}

export const getAppLetsEncryptEnabled = async (appName) => {
	return await appsClient.url("letsencrypt")
		.query({"name": appName})
		.get()
		.json(res => res["enabled"]);
}

export const addAppDomain = async (appName, domain) => {
	return await appsClient.url("domain")
		.post({"name": appName, "domain": domain})
		.res(res => res.ok);
}

export const removeAppDomain = async (appName, domain) => {
	return await appsClient.url("domain")
		.json({"name": appName, "domain": domain})
		.delete()
		.res(res => res.ok);
}

export const getAppLogs = async (appName) => {
	return await appsClient.url("logs")
		.query({"name": appName})
		.get()
		.json(res => res["logs"]);
}

export const getAppConfig = async (appName) => {
	return await appsClient.url("config")
		.query({"name": appName})
		.get()
		.json(res => res["config"]);
}

export const setAppConfig = async ({appName, newConfig}) => {
	return await appsClient.url("config")
		.post({"name": appName, "config": newConfig})
		.res(res => res.ok);
}

export const getAppStorages = async (appName) => {
	return await appsClient.url("storage")
		.query({"name": appName})
		.get()
		.json(res => res);
}

export const mountAppStorage = async (appName, options) => {
	return await appsClient.url("storage/mount")
		.post({"name": appName, ...options})
		.res(res => res.ok);
}

export const unmountAppStorage = async (appName, options) => {
	return await appsClient.url("storage/unmount")
		.post({"name": appName, ...options})
		.res(res => res.ok);
}

export const getSelectedBuilder = async (appName) => {
	return await appsClient.url("builder")
		.query({"name": appName})
		.get()
		.json(res => {
			return res["selected"];
		});
}

export const setSelectedBuilder = async (appName, builder) => {
	return await appsClient.url("builder")
		.post({"name": appName, "builder": builder})
		.res(res => res.ok);
}

export const getAppBuildDirectory = async (appName) => {
	return await appsClient.url("build-directory")
		.query({"name": appName})
		.get()
		.json(res => res["directory"]);
}

export const setAppBuildDirectory = async (appName, newDir) => {
	return await appsClient.url("build-directory")
		.post({"name": appName, "directory": newDir})
		.res(res => res.ok);
}

export const clearAppBuildDirectory = async (appName) => {
	return await appsClient.url("build-directory")
		.query({"name": appName})
		.delete()
		.res(res => res.ok);
}
