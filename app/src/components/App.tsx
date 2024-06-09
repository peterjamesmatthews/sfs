import { useAuth0 } from "@auth0/auth0-react";
import { useSelector } from "react-redux";
import { selectAccessToken } from "../store/slices/auth";
import Me from "./Me";
import SignIn from "./SignIn";
import SignUp from "./SignUp";

export default function App() {
  const accessToken = useSelector(selectAccessToken);
  const { user, isAuthenticated } = useAuth0();

  return (
    <>
      <SignUp />
      {accessToken ? <Me /> : <SignIn />}
      {isAuthenticated && user ? user.name : null}
    </>
  );
}
