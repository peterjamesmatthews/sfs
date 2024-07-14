import { useAuth0 } from "@auth0/auth0-react";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import store from "../store";
import { getTokensFromAuth0, selectAccessToken } from "../store/slices/auth";

export default function SignIn() {
  const accessToken = useSelector(selectAccessToken);
  const { loginWithRedirect, isAuthenticated, getAccessTokenSilently } =
    useAuth0();

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
    <button type="button" onClick={() => loginWithRedirect()}>
      Sign in with Auth0
    </button>
  );
}
