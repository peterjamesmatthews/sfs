import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./components/App.tsx";
import defaultTheme from "./themes/index.ts";

const root = document.getElementById("root");
if (root == null) throw new Error("no root");

const client = new ApolloClient({
  uri: "/graph",
  cache: new InMemoryCache(),
});

ReactDOM.createRoot(root).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <ThemeProvider theme={defaultTheme}>
        <App />
      </ThemeProvider>
    </ApolloProvider>
  </React.StrictMode>
);
