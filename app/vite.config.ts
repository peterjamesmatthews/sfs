import react from "@vitejs/plugin-react-swc";
import { defineConfig } from "vite";
import GQLCodegen from "vite-plugin-graphql-codegen";

// https://vitejs.dev/config/
export default defineConfig(() => {
	const {
		SERVER_HOSTNAME = "server",
		SERVER_WEB_PORT = 8080,
		SERVER_GRAPH_ENDPOINT = "graph",
	} = process.env;

	return {
		plugins: [react(), GQLCodegen()],
		server: {
			host: true,
			cors: true,
			proxy: {
				[`/${SERVER_GRAPH_ENDPOINT}`]: `http://${SERVER_HOSTNAME}:${SERVER_WEB_PORT}`,
			},
		},
	};
});
