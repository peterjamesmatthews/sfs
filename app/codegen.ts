import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
	schema: "http://server:8080/graphql",
	documents: ["src/**/*.gql"],
	ignoreNoDocuments: true,
	generates: {
		"src/gql/": {
			preset: "client",
			presetConfig: {
				gqlTagName: "gql",
			},
		},
	},
};

export default config;
