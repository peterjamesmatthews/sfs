import { ApolloClient, ApolloLink, InMemoryCache } from "@apollo/client";
import store from "../store";
import { selectAccessToken } from "../store/slices/auth";

export default new ApolloClient({
  uri: "/graph",
  cache: new InMemoryCache(),
  // TODO fix me
  link: new ApolloLink((operation, forward) => {
    const token = selectAccessToken(store.getState());
    if (!token) return forward(operation);
    operation.setContext({ headers: { Authorization: `Bearer ${token}` } });
    return forward(operation);
  }),
});
