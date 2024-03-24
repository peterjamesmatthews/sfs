import { gql } from "../generated";

export default gql(`
query GetNodeByURI($uri: String!) {
  getNodeByURI(uri: $uri) {
    id
    name
    owner { id }
    parent { id }
  }
}
`);
