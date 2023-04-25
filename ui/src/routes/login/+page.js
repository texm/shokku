import {readAuthCookie} from "$lib/auth";
import {redirect} from "@sveltejs/kit";

export const ssr = false;

export async function load({url}) {
	const authDetails = await readAuthCookie();
	if (authDetails) {
		throw redirect(307, "/");
	}

	return {};
}