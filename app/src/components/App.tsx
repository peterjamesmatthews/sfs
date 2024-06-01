import { useSelector } from "react-redux";
import { selectAccessToken } from "../store/slices/auth";
import Me from "./Me";
import SignIn from "./SignIn";
import SignUp from "./SignUp";

export default function App() {
  const accessToken = useSelector(selectAccessToken);

  return (
    <>
      <SignUp />
      {accessToken ? <Me /> : <SignIn />}
    </>
  );
}
