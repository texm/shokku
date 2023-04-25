import {getServiceType} from "$lib/api";
import {error} from '@sveltejs/kit';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	try {
		const serviceName = params.name;
		const serviceType = await getServiceType(serviceName);
		return {serviceName, serviceType};
	} catch (e) {
		throw error(404, `no service with name ${params.name}`)
	}
}