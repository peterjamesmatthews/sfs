import { gql } from "../generated";

export default gql(`
mutation RefreshTokens($refresh: String!) {
  refreshTokens(refresh: $refresh) {
    access
    refresh
  }
}
`);
