import { useQuery } from "@apollo/client";
import { Typography } from "@mui/material";
import Me from "../graphql/query/Me";

export default function App() {
  const { loading, error, data } = useQuery(Me);
  if (loading) return <Typography>Loading...</Typography>;
  if (error) return <Typography>Error {error.message}</Typography>;
  if (!data) return <Typography>No data</Typography>;
  return <Typography>Hello, {data.me.name}!</Typography>;
}
