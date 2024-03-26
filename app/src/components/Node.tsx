import { useQuery } from "@apollo/client";
import GetNodeByURI from "../graphql/query/GetNodeByURI";
import File from "./File";
import Folder from "./Folder";

type NodeProps = {
	uri: string;
};

export default function Node({ uri }: NodeProps) {
	const { loading, error, data } = useQuery(GetNodeByURI, {
		variables: { uri },
	});

	if (loading) return <>Loading node {uri}...</>;
	if (error)
		return (
			<>
				Error getting {uri}: {error.message}
			</>
		);
	if (!data?.getNodeByURI) return <>{uri} not found.</>;

	const node = data.getNodeByURI;

	switch (node.__typename) {
		case "Folder":
			return <Folder id={node.id} uri={uri} />;
		case "File":
			return <File id={node.id} uri={uri} />;
		default:
			return <>Unknown node type: {node.__typename}</>;
	}
}
