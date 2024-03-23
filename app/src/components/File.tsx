import { useQuery } from "@apollo/client";
import { gql } from "../gql";
import Typography from "@mui/material/Typography";

const GET_FILE_BY_ID = gql(`
query GetFileByID($id: ID!) {
  getFileById(id: $id) {
    name
    owner { id }
    parent { id }
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
				Error getting file {id}: {error.message}
			</>
		);
	if (!data?.getFileById) return <>File {id} not found</>;

	const file = data.getFileById;
	return <Typography>{file.name}</Typography>;
}
