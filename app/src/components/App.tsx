import { useSelector } from "react-redux";
import { RouterProvider } from "react-router-dom";
import router from "../router";
import { selectAccessToken } from "../store/slices/auth";
import SignIn from "./SignIn";

export default function App() {
  const accessToken = useSelector(selectAccessToken);
  if (!accessToken) return <SignIn />;
  return <RouterProvider router={router} />;
}
