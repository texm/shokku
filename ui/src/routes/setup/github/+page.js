import {getGithubAppSetupStatus} from "$lib/api";
import {redirect} from "@sveltejs/kit";

export async function load({ params }) {
	const status = await getGithubAppSetupStatus();
	if (status["created"]) throw redirect(302, "/setup/github/install");
	return {};
}