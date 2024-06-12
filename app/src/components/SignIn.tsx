import { useAuth0 } from "@auth0/auth0-react";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import store from "../store";
import {
  getTokensFromAuth0,
  getTokensFromNameAndPassword,
  selectAccessToken,
} from "../store/slices/auth";

export default function SignIn() {
  const accessToken = useSelector(selectAccessToken);
  const { loginWithRedirect, logout, isAuthenticated, getAccessTokenSilently } =
    useAuth0();

  const handleSubmit: React.FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const name = formData.get("name") as string;
    const password = formData.get("password") as string;
    store.dispatch(getTokensFromNameAndPassword({ name, password }));
  };

  // if a user is authenticated with Auth0 and not with the sfs, dispatch action
  // to get tokens from the sfs from the Auth0's token
  useEffect(() => {
    if (isAuthenticated && !accessToken) {
      getAccessTokenSilently().then((token) => {
        store.dispatch(getTokensFromAuth0({ token }));
      });
    }
  }, [isAuthenticated, accessToken, getAccessTokenSilently]);

  return (
    <form onSubmit={handleSubmit}>
      <input type="name" name="name" placeholder="Name" />
      <input type="password" name="password" placeholder="Password" />
      <button type="submit">Sign In</button>
      <button
        onClick={() =>
          loginWithRedirect({ authorizationParams: { connection: "github" } })
        }
      >
        Sign in with Auth0
      </button>
      <button onClick={() => logout()}>Sign out with Auth0</button>
    </form>
  );
}
