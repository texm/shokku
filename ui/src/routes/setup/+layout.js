import {redirect} from "@sveltejs/kit";
import {checkSetupKeySet} from "$lib/auth";

export const ssr = false;

export async function load({url}) {
	if (url.pathname !== "/setup" && !checkSetupKeySet()) {
		throw redirect(302, "/setup")
	}

	return {};
}