import {
  PayloadAction,
  createAsyncThunk,
  createSelector,
  createSlice,
} from "@reduxjs/toolkit";
import { AppDispatch, StoreState } from "..";
import apollo from "../../apollo";
import {
  CreateUserMutation,
  GetTokensQuery,
  Tokens,
  User,
} from "../../graphql/generated/graphql";
import CreateUser from "../../graphql/query/CreateUser";
import GetTokens from "../../graphql/query/GetTokens";

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
    gotTokens: (state, payload: PayloadAction<GetTokensQuery["getTokens"]>) => {
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

  const createdUser = data.createUser;

  console.debug(`created user ${createdUser.name} with id ${createdUser.id}`);

  return createdUser;
});

/** Gets a user's tokens from their name and password. */
export const getTokens = createAsyncThunk<
  GetTokensQuery["getTokens"],
  { name: string; password: string },
  { dispatch: AppDispatch; state: StoreState }
>("auth/getTokens", async ({ name, password }, { dispatch }) => {
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

  const { access, refresh } = data.getTokens;

  console.debug(`got tokens: access=${access}, refresh=${refresh}`);

  dispatch(auth.actions.gotTokens(data.getTokens));

  return data.getTokens;
});
