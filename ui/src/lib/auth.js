import Cookies from "js-cookie";
import jwt_decode from "jwt-decode";
import {refreshAuthToken} from "./api";

const AuthDataCookieKey = "auth_data";

const timeTillExpiry = (authCookie) => {
	const expiryTimestampMillis = authCookie["exp"] * 1000;
	const expiryTime = new Date(expiryTimestampMillis);
	return expiryTime.getTime() - Date.now();
}

export const readAuthCookie = () => {
	const val = Cookies.get(AuthDataCookieKey);
	if (!val) return null;

	const payload = jwt_decode(val);
	const timeLeft = timeTillExpiry(payload);
	if (timeLeft < 0) return null;
	return payload;
}

// check every 30 seconds, refresh token at 2 minutes left
const refreshTokenCutoffMs = 2 * 60_000;
const refreshIntervalDelayMs = 30_000;

const tryRefreshAuth = async () => {
	const auth = readAuthCookie();
	if (!auth) return;
	if (timeTillExpiry(auth) - refreshTokenCutoffMs > 0) return;

	try {
		await refreshAuthToken();
	} catch (e) {
		console.error("error refreshing auth", e);
	}
}

let refreshIntervalId;
export const startAuthRefresh = async () => {
	if (refreshIntervalId) stopAuthRefresh();
	refreshIntervalId = setInterval(tryRefreshAuth, refreshIntervalDelayMs)
}

export const stopAuthRefresh = () => {
	if (refreshIntervalId) clearInterval(refreshIntervalId);
	refreshIntervalId = null;
}

const generateRandomString = (n) => {
	const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
	let vals = new Uint8Array(n);
	window.crypto.getRandomValues(vals);
	return String.fromCharCode.apply(null, vals.map(
		x => chars.charCodeAt(x % chars.length)));
}

export const createStoredState = (storageKey) => {
	const state = generateRandomString(16);
	localStorage.setItem(storageKey, state);
	return state;
}

export const checkStoredState = (storageKey, supplied) => {
	const storedState = localStorage.getItem(storageKey);
	if (supplied === storedState) {
		localStorage.removeItem("github_state");
		return true;
	}
	return false;
}

export const getSetupStatus = async (refresh) => {
	if (!refresh) {
		const storedStatus = localStorage.getItem("setup_status");
		if (!!storedStatus) return JSON.parse(storedStatus);
	}

	const status = await fetch("/api/setup/status").then(r => r.json());
	setSetupStatus(status)
	return status;
}

export const setSetupStatus = (status) => {
	localStorage.setItem("setup_status", JSON.stringify(status));
}

export const checkSetupKeySet = () => {
	const storedKey = localStorage.getItem("setup_key");
	return !!storedKey;
}

export const setSetupKey = (key) => {
	localStorage.setItem("setup_key", key);
}