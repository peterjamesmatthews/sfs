import React from "react";
import ReactDOM from "react-dom/client";
import Root from "./components/Root.tsx";
import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";

const root = document.getElementById("root");
if (root == null) throw new Error("no root");

const client = new ApolloClient({
	uri: "/graphql",
	cache: new InMemoryCache(),
	headers: { Authorization: "Nick" },
});

ReactDOM.createRoot(root).render(
	<React.StrictMode>
		<ApolloProvider client={client}>
			<Root />
		</ApolloProvider>
	</React.StrictMode>,
);
