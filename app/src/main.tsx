import { ApolloProvider } from "@apollo/client";
import { Auth0Provider } from "@auth0/auth0-react";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import React from "react";
import ReactDOM from "react-dom/client";
import { Provider as ReduxProvider } from "react-redux";
import client from "./apollo";
import App from "./components/App";
import store from "./store";
import theme from "./theme";

const root = document.getElementById("root");
if (root == null) throw new Error("no root");

ReactDOM.createRoot(root).render(
  <React.StrictMode>
    <Auth0Provider
      domain={import.meta.env.VITE_AUTH0_DOMAIN}
      clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
      authorizationParams={{ redirect_uri: window.location.origin }}
    >
      <ReduxProvider store={store}>
        <ApolloProvider client={client}>
          <ThemeProvider theme={theme}>
            <App />
          </ThemeProvider>
        </ApolloProvider>
      </ReduxProvider>
    </Auth0Provider>
  </React.StrictMode>
);
