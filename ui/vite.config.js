import path from "path";
import { sveltekit } from '@sveltejs/kit/vite';

const config = {
	plugins: [sveltekit()],
	clearScreen: false,
	logLevel: "info",
	server: {
		proxy: {
			"/api/exec/socket": {
				target: "ws://localhost:5330",
				changeOrigin: true,
				ws: true,
			},
			"/api": {
				target: "http://localhost:5330",
				changeOrigin: true,
			}
		}
	},
	resolve: {
		alias: {
			$src: path.resolve("./src"),
			$components: path.resolve("./src/components"),
			$common: path.resolve("./src/components/common"),
		}
	},
}

export default config;