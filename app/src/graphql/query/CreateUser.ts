import { gql } from "../generated";

export default gql(`
mutation CreateUser($name: String!, $password: String!) {
  createUser(name: $name, password: $password) {
    id
    name
  }
}
`);
