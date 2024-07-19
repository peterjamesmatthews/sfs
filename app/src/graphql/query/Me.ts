import { gql } from "../generated";

export default gql(`
query Me {
  me {
    email
  }
}
`);
