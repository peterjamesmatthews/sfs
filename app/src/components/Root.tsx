import { useQuery } from "@apollo/client";
import { gql } from "../gql/gql";
import File from "./File";
import Folder from "./Folder";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListSubheader from "@mui/material/ListSubheader";

const GET_ROOT = gql(`
query GetRoot {
  getRoot {
    children {
      id
    }
  }
}
`);

export default function Root() {
	const { loading, error, data } = useQuery(GET_ROOT);

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
