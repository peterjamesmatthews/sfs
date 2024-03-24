import { gql } from "../generated";

export default gql(`
query GetRoot {
  getRoot {
    children {
      id
    }
  }
}
`);
