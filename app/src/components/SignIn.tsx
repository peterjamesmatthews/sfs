import { useAuth0 } from "@auth0/auth0-react";
import store from "../store";
import { getTokens } from "../store/slices/auth";

export default function SignIn() {
  const { loginWithRedirect, logout } = useAuth0();

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
      <button onClick={() => loginWithRedirect()}>Sign in with Auth0</button>
      <button onClick={() => logout()}>Sign out with Auth0</button>
    </form>
  );
}
