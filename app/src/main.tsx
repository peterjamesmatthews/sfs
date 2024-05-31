import { ApolloProvider } from "@apollo/client";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import React from "react";
import ReactDOM from "react-dom/client";
import { Provider as ReduxProvider } from "react-redux";
import client from "./apollo";
import SignUp from "./components/SignUp.tsx";
import store from "./store";
import theme from "./theme";

const root = document.getElementById("root");
if (root == null) throw new Error("no root");

ReactDOM.createRoot(root).render(
  <React.StrictMode>
    <ReduxProvider store={store}>
      <ApolloProvider client={client}>
        <ThemeProvider theme={theme}>
          <SignUp />
        </ThemeProvider>
      </ApolloProvider>
    </ReduxProvider>
  </React.StrictMode>
);
