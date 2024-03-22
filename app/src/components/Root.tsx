import { gql, useQuery } from "@apollo/client";

const GET_ROOT = gql`
query Node {
  node {
    ... on Folder {
      children {
        owner {
          id
          name
        }
      }
    }
  }
}
`;

export default function Root() {
	const getRoot = useQuery(GET_ROOT);

	if (getRoot.loading) return <>Loading...</>;
	if (getRoot.error) return <>Error: {getRoot.error.message}</>;

	return <>{JSON.stringify(getRoot.data)}</>;
}
