import { jwtDecode } from "jwt-decode";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import { RouterProvider } from "react-router-dom";
import router from "../router";
import store from "../store";
import { refreshTokens, selectAccessToken } from "../store/slices/auth";
import SignIn from "./SignIn";

export default function App() {
  const accessToken = useSelector(selectAccessToken);

  useEffect(() => {
    if (!accessToken) return;
    const { exp } = jwtDecode(accessToken);
    if (!exp) return;
    store.dispatch(refreshTokens({ expiresAt: new Date(exp * 1000) }));
  }, [accessToken]);

  if (!accessToken) return <SignIn />;
  return <RouterProvider router={router} />;
}
