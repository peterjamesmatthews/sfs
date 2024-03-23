import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";
import defaultTheme from "./themes/index.ts";
import router from "./router/index.tsx";

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
			<ThemeProvider theme={defaultTheme}>
				<RouterProvider router={router} />
			</ThemeProvider>
		</ApolloProvider>
	</React.StrictMode>,
);
