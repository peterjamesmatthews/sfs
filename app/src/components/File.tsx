import { useQuery } from "@apollo/client";
import Typography from "@mui/material/Typography";
import GetFileById from "../graphql/query/GetFileById";

type FileProps = {
	id: string;
};

export default function File({ id }: FileProps) {
	const { loading, error, data } = useQuery(GetFileById, {
		variables: { id },
	});

	if (loading) return <>Loading file {id}...</>;
	if (error)
		return (
			<>
				Error getting file {id}: {error.message}
			</>
		);
	if (!data?.getFileById) return <>File {id} not found</>;

	const file = data.getFileById;
	return <Typography>{file.name}</Typography>;
}
