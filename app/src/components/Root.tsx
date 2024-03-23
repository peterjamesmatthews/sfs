import { useQuery } from "@apollo/client";
import File from "./File";
import Folder from "./Folder";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListSubheader from "@mui/material/ListSubheader";
import GetRoot from "../graphql/query/GetRoot";

export default function Root() {
	const { loading, error, data } = useQuery(GetRoot);

	if (loading) return <>Loading root...</>;
	if (error) return <>Error getting root: {error.message}</>;
	if (!data?.getRoot) return <>Root not found</>;

	const children = data.getRoot.children;
	return (
		<>
			<List subheader={<ListSubheader>/</ListSubheader>}>
				{children.map((child) => (
					<ListItem key={child.id}>
						{child.__typename === "Folder" ? (
							<Folder id={child.id} />
						) : (
							<File id={child.id} />
						)}
					</ListItem>
				))}
			</List>
		</>
	);
}
