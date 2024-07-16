import { gql } from "../generated";

export default gql(`
query GetNodeFromPath($path: String!) {
  getNodeFromPath(path: $path) {
    id
    name
    parent {
      id
      name
    }
    ...on Folder {
      children {
        id
        name
      }
    }
    ... on File {
      content
    }
  }  
}
`);
