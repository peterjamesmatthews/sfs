import { gql } from "../generated";

export default gql(`
query GetFileByID($id: ID!) {
  getFileById(id: $id) {
    name
    owner { id }
    parent { id }
    content
  }
}
`);
