import {createApiClient} from "./client.js";

const execClient = createApiClient("/exec/");

export const getExecutionStatus = async (id) => {
	return await execClient.url("status")
		.query({"execution_id": id})
		.get()
		.json(res => res);
}

export const executeCommandInProcess = async (appName, processName, command) => {
	return await execClient.url("command")
		.post({appName, processName, command})
		.json(res => res);
}
/*
const getSocketURI = (appName, processName) => {
	const loc = window.location;
	const proto = (loc.protocol === 'https:') ? "wss:" : "ws:";
	const queryParams = `app_name=${appName}&process_name=${processName}`
	return `${proto}//${loc.host}/api/exec/socket?${queryParams}`
}

export const getAppProcessExecutionSocket = (appName, processName) => {
	const uri = getSocketURI(appName, processName);
	return new WebSocket(uri)
}
*/