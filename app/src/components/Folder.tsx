import { useQuery } from "@apollo/client";
import Typography from "@mui/material/Typography";
import GetFolderById from "../graphql/query/GetFolderById";

type FolderProps = {
	id: string;
};

export default function Folder({ id }: FolderProps) {
	const { loading, error, data } = useQuery(GetFolderById, {
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
	return <Typography>{folder.name}</Typography>;
}
