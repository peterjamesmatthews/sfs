import { gql } from "../generated";

export default gql(`
query GetTokensFromAuth0($token: String!) {
  getTokensFromAuth0(token: $token) {
    access
    refresh
  }
}
`);
