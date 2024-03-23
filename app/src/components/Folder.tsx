import { useQuery } from "@apollo/client";
import { gql } from "../gql";

const GET_FOLDER = gql(`
query GetFolderByID($id: ID!) {
  getFolderById(id: $id) {
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
				Error getting folder {id}: {error.message}
			</>
		);
	if (!data?.getFolderById) return <>Folder {id} not found</>;

	const folder = data.getFolderById;
	return <>{folder.name}</>;
}
