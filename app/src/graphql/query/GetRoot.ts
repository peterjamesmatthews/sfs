import { gql } from "../generated/gql";

export default gql(`
query GetRoot {
  getRoot {
    children {
      id
    }
  }
}
`);
