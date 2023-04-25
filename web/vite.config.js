import path from "path";
import { sveltekit } from '@sveltejs/kit/vite';

const config = {
	plugins: [sveltekit()],
	resolve: {
		alias: {
			$src: path.resolve("./src"),
			$components: path.resolve("./src/components"),
		}
	},
};

export default config;
