import Node from "./Node";
import URI from "./URI";

export default function App() {
	const uri = window.location.pathname;

	return (
		<>
			<URI uri={uri} />
			<Node uri={uri} />
		</>
	);
}
