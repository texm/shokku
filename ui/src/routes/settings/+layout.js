export async function load({ url }) {
	const currentPage = url.pathname.substring("/settings".length + 1);
	return {currentPage}
}