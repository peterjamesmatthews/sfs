import { gql } from "../generated";

export default gql(`
query GetTokens($name: String!, $password: String!) {
  getTokens(name: $name, password: $password) {
    access
    refresh
  }
}
`);
