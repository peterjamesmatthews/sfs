import { useQuery } from "@apollo/client";
import { gql } from "../gql/gql";
import File from "./File";
import Folder from "./Folder";

const GET_ROOT = gql(`
query Node {
  node {
    ... on Folder {
      children {
        id
      }
    }
  }
}
`);

export default function Root() {
	const getRoot = useQuery(GET_ROOT);

	if (getRoot.loading) return <>Loading...</>;
	if (getRoot.error) return <>Error: {getRoot.error.message}</>;
	if (getRoot.data === undefined) return <>No data</>;
	if (getRoot.data.node === undefined || getRoot.data.node === null)
		return <>No root</>;
	if (getRoot.data.node.__typename !== "Folder")
		return <>Root is not a folder</>;

	const children = getRoot.data.node.children;

	return (
		<>
			<ul>
				{children.map((child) => (
					<li key={child.id}>
						{child.__typename === "Folder" ? (
							<Folder id={child.id} />
						) : (
							<File id={child.id} />
						)}
					</li>
				))}
			</ul>
		</>
	);
}
