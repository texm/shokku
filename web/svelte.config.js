import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter(),
		paths: {
			base: process.env.NODE_ENV === "production" ? process.env.BASE_PATH : "",
			relative: false,
		},
	},
	preprocess: vitePreprocess(),
};

export default config;
