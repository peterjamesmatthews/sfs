import { useSelector } from "react-redux";
import store from "../store";
import { getTokens, selectAccessToken } from "../store/slices/auth";

export default function SignIn() {
  const accessToken = useSelector(selectAccessToken);
  if (accessToken) return <p>Signed in.</p>;

  const handleSubmit: React.FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const name = formData.get("name") as string;
    const password = formData.get("password") as string;
    store.dispatch(getTokens({ name, password }));
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="name" name="name" placeholder="Name" />
      <input type="password" name="password" placeholder="Password" />
      <button type="submit">Sign In</button>
    </form>
  );
}
