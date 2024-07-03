import { useAuth0 } from "@auth0/auth0-react";
import { useSelector } from "react-redux";
import { selectAccessToken } from "../store/slices/auth";

export default function SignOut() {
  const accessToken = useSelector(selectAccessToken);
  const { logout } = useAuth0();
  return accessToken === undefined ? null : (
    <button onClick={() => logout()}>Sign out with Auth0</button>
  );
}
