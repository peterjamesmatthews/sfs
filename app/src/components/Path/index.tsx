import { Typography } from "@mui/material";
import { Link, useLocation } from "react-router-dom";
import "./Path.css";

export default function Path() {
  const location = useLocation();

  /** Non-empty path segments. */
  const segments = location.pathname.split("/").filter(Boolean);

  /** Absolute paths for each segment. */
  const paths = segments.map(
    (_, i) => `/${segments.slice(0, i + 1).join("/")}`
  );

  /** Links to each path segment. */
  let links = paths.map((path, i) => (
    <Link className="Link" to={path}>
      /{segments[i]}
    </Link>
  ));

  return <Typography className="Path">{links}</Typography>;
}
