import { gql } from "../generated/gql";

export default gql(`
query GetFileByID($id: ID!) {
  getFileById(id: $id) {
    name
    owner { id }
    parent { id }
  }
}
`);
