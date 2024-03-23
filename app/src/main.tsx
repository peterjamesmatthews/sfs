import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import Typography from "@mui/material/Typography";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import React from "react";
import ReactDOM from "react-dom/client";
import { Link, RouterProvider, createBrowserRouter } from "react-router-dom";
import Root from "./components/Root.tsx";
import defaultTheme from "./themes/index.ts";

const root = document.getElementById("root");
if (root == null) throw new Error("no root");

const client = new ApolloClient({
	uri: "/graphql",
	cache: new InMemoryCache(),
	headers: { Authorization: "Nick" },
});

const router = createBrowserRouter([
	{
		path: "/",
		element: <Root />,
		errorElement: (
			<Typography sx={{ textDecoration: "none" }}>
				404 Not Found (<Link to="/">Return to Root</Link>)
			</Typography>
		),
	},
]);

ReactDOM.createRoot(root).render(
	<React.StrictMode>
		<ApolloProvider client={client}>
			<ThemeProvider theme={defaultTheme}>
				<RouterProvider router={router} />
			</ThemeProvider>
		</ApolloProvider>
	</React.StrictMode>,
);
