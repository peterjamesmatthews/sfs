import type { CodegenConfig } from "@graphql-codegen/cli";

const {
	SERVER_HOSTNAME = "localhost",
	SERVER_WEB_PORT = 8080,
	SERVER_GRAPH_ENDPOINT = "graph",
} = process.env;

const config: CodegenConfig = {
	schema: `http://${SERVER_HOSTNAME}:${SERVER_WEB_PORT}/${SERVER_GRAPH_ENDPOINT}`,
	documents: ["src/**/*.{ts,tsx}"],
	ignoreNoDocuments: true,
	generates: {
		"src/graphql/generated/": {
			preset: "client",
			presetConfig: { gqlTagName: "gql" },
		},
	},
};

export default config;
