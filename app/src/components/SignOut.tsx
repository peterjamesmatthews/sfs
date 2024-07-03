import { useAuth0 } from "@auth0/auth0-react";
import { useDispatch, useSelector } from "react-redux";
import auth, { selectAccessToken } from "../store/slices/auth";

export default function SignOut() {
  const dispatch = useDispatch();
  const accessToken = useSelector(selectAccessToken);
  const { logout } = useAuth0();

  const handleSignOut = () => {
    logout();
    dispatch(auth.actions.signOut());
  };

  return accessToken === undefined ? null : (
    <button onClick={handleSignOut}>Sign out with Auth0</button>
  );
}
