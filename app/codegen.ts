import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
	schema: "http://server:8080/graphql",
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
