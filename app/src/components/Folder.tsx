import { useQuery } from "@apollo/client";
import { gql } from "../gql";

const GET_FOLDER = gql(`
query FolderByID($id: ID) {
  node(id: $id) {
    ... on Folder {
      name
      owner { id }
      parent { id }
      children { id }
    }
  }
}
`);

type FolderProps = {
	id: string;
};

export default function Folder({ id }: FolderProps) {
	const { loading, error, data } = useQuery(GET_FOLDER, {
		variables: { id },
	});

	if (loading) return <>Loading folder {id}...</>;
	if (error)
		return (
			<>
				Error folder {id}: {error.message}
			</>
		);
	if (data === undefined || !data.node) return <>Folder {id} not found</>;
	if (data.node.__typename !== "Folder") return <>Node {id} is not a folder</>;

	const folder = data.node;
	return <>{folder.name}</>;
}
