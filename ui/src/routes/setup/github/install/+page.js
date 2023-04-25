import {getGithubAppInstallInfo} from "$lib/api";
import {redirect} from "@sveltejs/kit";

export const ssr = false;

export async function load() {
	const info = await getGithubAppInstallInfo();
	const url = info["install_url"];
	throw redirect(307, url);
}