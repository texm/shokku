import {createApiClient} from "./client.js";

const servicesClient = createApiClient("/services/");

export const createService = async (name, type, config) => {
	return await servicesClient.url(`create`)
		.post({name, config, type})
		.res(res => res.ok);
}

export const cloneService = async (name, newName) => {
	return await servicesClient.url(`clone`)
		.post({name, newName})
		.res(res => res.ok);
}

export const listServices = async () => {
	return await servicesClient.url("list")
		.get()
		.json(res => res["services"]);
}

export const getServiceType = async (serviceName) => {
	return await servicesClient.url("type")
		.query({"name": serviceName})
		.get()
		.json(res => res["type"]);
}

export const getServiceInfo = async (serviceName, serviceType) => {
	return await servicesClient.url("info")
		.query({"name": serviceName, "type": serviceType})
		.get()
		.json(res => res["info"]);
}

export const linkServiceToApp = async (serviceName, appName, options) => {
	return await servicesClient.url("link")
		.post({
			"service_name": serviceName,
			"app_name": appName,
			...options,
		})
		.json(res => res["execution_id"]);
}

export const unlinkServiceFromApp = async (serviceName, appName) => {
	return await servicesClient.url("unlink")
		.post({
			"service_name": serviceName,
			"app_name": appName,
		})
		.json(res => res["execution_id"]);
}

export const getServiceLinkedApps = async (serviceName, serviceType) => {
	return await servicesClient.url("linked-apps")
		.query({"name": serviceName, "type": serviceType})
		.get()
		.json(res => res["apps"]);
}

export const getServiceLogs = async (serviceName, serviceType) => {
	return await servicesClient.url("logs")
		.query({"name": serviceName, "type": serviceType})
		.get()
		.json(res => res["logs"]);
}

export const startService = async (serviceType, serviceName) => {
	return await servicesClient.url("start")
		.post({"name": serviceName, "type": serviceType})
		.res(res => res.ok);
}

export const stopService = async (serviceType, serviceName) => {
	return await servicesClient.url("stop")
		.post({"name": serviceName, "type": serviceType})
		.res(res => res.ok);
}

export const restartService = async (serviceType, serviceName) => {
	return await servicesClient.url("restart")
		.post({"name": serviceName, "type": serviceType})
		.res(res => res.ok);
}

export const destroyService = async (serviceType, serviceName) => {
	return await servicesClient.url("destroy")
		.post({"name": serviceName, "type": serviceType})
		.res(res => res.ok);
}

export const getServiceBackupReport = async (serviceName) => {
	return await servicesClient.url("backups/report")
		.query({"name": serviceName})
		.get()
		.json(res => res["report"]);
}

export const doServiceBackup = async (serviceName) => {
	return await servicesClient.url("backups/run")
		.post({"name": serviceName})
		.res(res => res.ok);
}

export const updateServiceBackupAuth = async (serviceName, config) => {
	return await servicesClient.url("backups/auth")
		.post({"name": serviceName, "config": config})
		.res(res => res.ok);
}

export const setServiceBackupsSchedule = async (serviceName, schedule) => {
	return await servicesClient.url("backups/schedule")
		.post({"name": serviceName, "schedule": schedule})
		.res(res => res.ok);
}

export const unscheduleServiceBackups = async (serviceName) => {
	return await servicesClient.url("backups/schedule")
		.body({"name": serviceName})
		.delete()
		.res(res => res.ok);
}

export const setServiceBackupEncryption = async (serviceName, passphrase) => {
	return await servicesClient.url("backups/encryption")
		.post({"name": serviceName, "passphrase": passphrase})
		.res(res => res.ok);
}

export const unsetServiceBackupEncryption = async (serviceName) => {
	return await servicesClient.url("backups/encryption")
		.body({"name": serviceName})
		.delete()
		.res(res => res.ok);
}