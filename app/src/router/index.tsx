import { Link, createBrowserRouter } from "react-router-dom";
import Root from "../components/Root";
import { Typography } from "@mui/material";

export default createBrowserRouter([
	{
		path: "/",
		element: <Root />,
		errorElement: (
			<Typography>
				404 Not Found (<Link to="/">Return to root</Link>)
			</Typography>
		),
	},
]);
