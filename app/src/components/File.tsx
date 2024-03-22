import { useQuery } from "@apollo/client";
import { gql } from "../gql";

const GET_FILE_BY_ID = gql(`
query FileByID($id: ID) {
  node(id: $id) {
    ... on File {
      name
      owner { id }
      parent { id }
      content
    }
  }
}
`);

type FileProps = {
	id: string;
};

export default function File({ id }: FileProps) {
	const { loading, error, data } = useQuery(GET_FILE_BY_ID, {
		variables: { id },
	});

	if (loading) return <>Loading file {id}...</>;
	if (error)
		return (
			<>
				Error file {id}: {error.message}
			</>
		);
	if (data === undefined || !data.node) return <>File {id} not found</>;
	if (data.node.__typename !== "File") return <>Node {id} is not a file</>;

	const file = data.node;
	return <>{file.name}</>;
}
