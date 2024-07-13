import { useQuery } from "@apollo/client";
import MeQuery from "../graphql/query/Me";

export default function Me() {
  const { loading, data, error } = useQuery(MeQuery);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;
  if (!data) return <div>No data</div>;

  return (
    <div>
      <h1>Hello {data.me.name}</h1>
    </div>
  );
}
