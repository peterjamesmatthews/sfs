import { PayloadAction, createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { AppDispatch, StoreState } from "..";
import apollo from "../../apollo";
import { User } from "../../graphql/generated/graphql";
import CreateUser from "../../graphql/query/CreateUser";

type AuthState = {
  user?: User;
};

const initialState: AuthState = {};

export default createSlice({
  name: "auth",
  initialState,
  reducers: {
    setUser: (state, payload: PayloadAction<User>) => {
      state.user = payload.payload;
    },
  },
});

/**
 * Creates a new user.
 */
export const createUser = createAsyncThunk<
  void,
  {
    /** The name of the user to create. */
    name: string;
    /** The password the created user will sign in with. */
    password: string;
  },
  { dispatch: AppDispatch; state: StoreState }
>("auth/createUser", async ({ name, password }) => {
  // send mutation to create user
  const { data, errors } = await apollo.mutate({
    mutation: CreateUser,
    variables: { name, password },
  });

  // throw any errors
  if (errors)
    throw new Error(
      "failed to create user: " + errors.map((e) => e.message).join(", ")
    );
  else if (!data?.createUser) throw new Error("failed to create user"); // unexpected

  /** The name and id of the created user. */
  const createdUser = data.createUser;

  console.log(`created user ${createdUser.name} with id ${createdUser.id}`);
});
