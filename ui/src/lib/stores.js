import {readable, writable} from "svelte/store";
import {getExecutionStatus} from "./api.js";

const createAppThemeStore = () => {
	const themeKey = "app_theme";
	const defaultTheme = "business";
	const storedValue = localStorage.getItem(themeKey);
	const {subscribe, set} = writable(storedValue || defaultTheme);
	return {
		subscribe,
		set: (val) => {
			localStorage.setItem(themeKey, val);
			set(val);
		}
	}
}

let completionCallbacks = {}
const createCommandExecutionIdsStore = () => {
	const {subscribe, update} = writable([]);
	const remove = (id) => update(val => val.filter(el => el !== id));
	return {
		subscribe,
		addID: (id) => {
			update(ids => [...ids, id]);
			return new Promise(res => completionCallbacks[id] = res);
		},
		signalFinished: (id, success) => {
			remove(id);
			if (completionCallbacks[id]) {
				completionCallbacks[id](success);
				delete completionCallbacks[id];
			}
		},
		remove,
	}
}

const pollCommandExecutionsFn = (idsStore) => {
	const pollIntervalMs = 5000;

	let ids = [];
	idsStore.subscribe(val => ids = val);

	return (set) => {
		let isPolling = {}
		const interval = setInterval(async () => {
			let didPoll = false;

			let statuses = {}
			for (const i in ids) {
				const id = ids[i];

				if (isPolling[id]) return;
				if (!completionCallbacks[id]) return;
				didPoll = true;

				isPolling[id] = true;
				try {
					let status = await getExecutionStatus(id);
					statuses[id] = status;
					isPolling[id] = false

					if (status["finished"]) {
						idsStore.signalFinished(id, status["success"]);
					}
				} catch (e) {
					if (e.status === 404) {
						delete isPolling[id];
						idsStore.remove(id);
						return;
					}
				}
				isPolling[id] = false;
			}

			if (didPoll) set(statuses);
		}, pollIntervalMs);

		return () => clearInterval(interval);
	}
}

export const appTheme = createAppThemeStore();
export const commandExecutionIds = createCommandExecutionIdsStore();
export const commandExecutions = readable({}, pollCommandExecutionsFn(commandExecutionIds));
export const executionIdDescriptions = writable({});