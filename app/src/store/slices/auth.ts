import {
  PayloadAction,
  createAsyncThunk,
  createSelector,
  createSlice,
} from "@reduxjs/toolkit";
import { jwtDecode } from "jwt-decode";
import store, { AppDispatch, StoreState } from "..";
import apollo from "../../apollo";
import {
  CreateUserMutation,
  GetTokensQuery,
  RefreshTokensMutation,
  Tokens,
  User,
} from "../../graphql/generated/graphql";
import CreateUser from "../../graphql/query/CreateUser";
import GetTokens from "../../graphql/query/GetTokens";
import RefreshTokens from "../../graphql/query/RefreshTokens";

type AuthState = {
  user?: User;
  tokens?: Tokens;
};

const initialState: AuthState = {};

const auth = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setUser: (state, payload: PayloadAction<User>) => {
      state.user = payload.payload;
    },
    gotTokens: (state, payload: PayloadAction<Tokens>) => {
      state.tokens = payload.payload;
    },
  },
});

export default auth;

export function selectTokens(state: StoreState) {
  return state.auth.tokens;
}

export const selectAccessToken = createSelector(
  selectTokens,
  (tokens) => tokens?.access
);

export const selectRefreshToken = createSelector(
  selectTokens,
  (tokens) => tokens?.refresh
);

/** Creates a new user. */
export const createUser = createAsyncThunk<
  CreateUserMutation["createUser"],
  {
    /** The name of the user to create. */
    name: string;
    /** The password the created user will sign in with. */
    password: string;
  },
  { dispatch: AppDispatch; state: StoreState }
>("auth/createUser", async ({ name, password }) => {
  const { data, errors } = await apollo.mutate({
    mutation: CreateUser,
    variables: { name, password },
  });

  if (errors) {
    const error = new Error(
      `failed to create user: ${errors.map((e) => e.message).join(", ")}`
    );
    console.error(error);
    throw error;
  } else if (!data?.createUser) throw new Error("failed to create user"); // unexpected

  return data.createUser;
});

/** Gets a user's tokens from their name and password. */
export const getTokens = createAsyncThunk<
  GetTokensQuery["getTokens"],
  { name: string; password: string },
  { dispatch: AppDispatch }
>("auth/getTokens", async ({ name, password }, { dispatch }) => {
  // query for tokens
  const { data, errors } = await apollo.query({
    query: GetTokens,
    variables: { name, password },
  });
  if (errors) {
    const error = new Error(
      "failed to get tokens: " + errors.map((e) => e.message).join(", ")
    );
    console.error(error);
    throw error;
  } else if (!data?.getTokens) throw new Error("failed to get tokens");

  // store the tokens in the store
  dispatch(auth.actions.gotTokens(data.getTokens));

  // decode access token
  const accessToken = jwtDecode(data.getTokens.access);

  // schedule refresh of tokens
  if (accessToken.exp)
    dispatch(refreshTokens({ expiresAt: new Date(accessToken.exp * 1000) }));
  else console.warn("access token has no expiration date");

  return data.getTokens;
});

/**
 * refreshTokens is an async thunk that refresh the user's access tokens 10
 * seconds before they expire.
 */
export const refreshTokens = createAsyncThunk<
  RefreshTokensMutation["refreshTokens"],
  { expiresAt: Date },
  { dispatch: AppDispatch }
>("auth/refreshTokens", async ({ expiresAt }, { dispatch }) => {
  // calculate the amount of milliseconds until 10 seconds before the tokens expire
  const refreshAt = new Date(expiresAt.getTime() - 10 * 1000);

  // wait until it's time to refresh the tokens
  await new Promise((r) => setTimeout(r, refreshAt.getTime() - Date.now()));

  // get the current refresh token
  const refreshToken = selectRefreshToken(store.getState());
  if (!refreshToken) throw new Error("no refresh token");

  // refresh the tokens
  const { data, errors } = await apollo.mutate({
    mutation: RefreshTokens,
    variables: { refresh: refreshToken },
  });
  if (errors) {
    const error = new Error(
      "failed to refresh tokens: " + errors.map((e) => e.message).join(", ")
    );
    console.error(error);
    throw error;
  } else if (!data?.refreshTokens) throw new Error("failed to refresh tokens");

  // store the new tokens in the store
  dispatch(auth.actions.gotTokens(data.refreshTokens));

  // decode access token
  const accessToken = jwtDecode(data.refreshTokens.access);

  // schedule refresh of tokens
  if (accessToken.exp)
    dispatch(refreshTokens({ expiresAt: new Date(accessToken.exp * 1000) }));
  else console.warn("access token has no expiration date");

  return data.refreshTokens;
});
