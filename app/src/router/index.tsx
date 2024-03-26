import { createBrowserRouter } from "react-router-dom";
import App from "../components/App";

export default createBrowserRouter([
	{
		path: "*",
		element: <App />,
	},
]);
