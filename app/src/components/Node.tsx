import { useQuery } from "@apollo/client";
import { useLocation } from "react-router-dom";
import GetNodeFromPath from "../graphql/query/GetNodeFromPath";

export default function Node() {
  const location = useLocation();
  const path = location.pathname;
  const { data, loading, error } = useQuery(GetNodeFromPath, {
    variables: { path },
  });

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;
  if (!data) return <p>No data</p>;
  return "Node";
}
