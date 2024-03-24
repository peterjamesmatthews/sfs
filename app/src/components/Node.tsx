import { useQuery } from "@apollo/client";
import GetNodeByURI from "../graphql/query/GetNodeByURI";

export default function Node() {
	const uri = window.location.pathname;

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
	return <>Node: {node.name}</>;
}
