import { ApolloClient, InMemoryCache } from "@apollo/client";

export default new ApolloClient({
  uri: "/graph",
  cache: new InMemoryCache(),
});
