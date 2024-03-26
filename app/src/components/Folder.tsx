import { useQuery } from "@apollo/client";
import Typography from "@mui/material/Typography";
import GetFolderById from "../graphql/query/GetFolderById";
import List from "@mui/material/List";

type FolderProps = {
	id: string;
	uri: string;
};

export default function Folder({ id, uri }: FolderProps) {
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
	return (
		<List>
			{folder.children.map((child) => {
				return (
					<a
						href={`${uri}${child.name}${
							child.__typename === "Folder" ? "/" : ""
						}`}
					>
						<Typography key={child.id}>{child.name}</Typography>
					</a>
				);
			})}
		</List>
	);
}
