import { useSelector } from "react-redux";
import { selectAccessToken } from "../store/slices/auth";
import Me from "./Me";
import SignIn from "./SignIn";
import SignOut from "./SignOut";

export default function App() {
  const accessToken = useSelector(selectAccessToken);

  return (
    <>
      {accessToken ? (
        <>
          <Me />
          <SignOut />
        </>
      ) : (
        <SignIn />
      )}
    </>
  );
}
