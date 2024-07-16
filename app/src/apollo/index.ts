import {
  ApolloClient,
  ApolloLink,
  HttpLink,
  InMemoryCache,
} from "@apollo/client";
import store from "../store";
import { selectAccessToken } from "../store/slices/auth";

const link = ApolloLink.from([
  new ApolloLink((operation, forward) => {
    const token = selectAccessToken(store.getState());
    if (token)
      operation.setContext({ headers: { Authorization: `Bearer ${token}` } });
    return forward(operation);
  }),
  new HttpLink({ uri: "/graph" }),
]);

export default new ApolloClient({
  cache: new InMemoryCache(),
  link,
  connectToDevTools: true,
});
