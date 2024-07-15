import { ApolloProvider } from "@apollo/client";
import { Auth0Provider } from "@auth0/auth0-react";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import React from "react";
import ReactDOM from "react-dom/client";
import { Provider as ReduxProvider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import client from "./apollo";
import App from "./components/App";
import store, { persistor } from "./store";
import theme from "./theme";

const root = document.getElementById("root");
if (root == null) throw new Error("missing root");

ReactDOM.createRoot(root).render(
  <React.StrictMode>
    <Auth0Provider
      domain={import.meta.env.VITE_AUTH0_DOMAIN}
      clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
      authorizationParams={{ redirect_uri: window.location.origin }}
    >
      <ReduxProvider store={store}>
        <PersistGate loading={null} persistor={persistor}>
          <ApolloProvider client={client}>
            <ThemeProvider theme={theme}>
              <App />
            </ThemeProvider>
          </ApolloProvider>
        </PersistGate>
      </ReduxProvider>
    </Auth0Provider>
  </React.StrictMode>,
);
