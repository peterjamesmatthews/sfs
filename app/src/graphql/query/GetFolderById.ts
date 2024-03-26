import { gql } from "../generated";

export default gql(`
query GetFolderByID($id: ID!) {
  getFolderById(id: $id) {
    ... on Folder {
      name
      owner { id }
      parent { id }
      children { id name }
    }
  }
}
`);
