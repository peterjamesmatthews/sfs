import {
	type PayloadAction,
	createAsyncThunk,
	createSelector,
	createSlice,
} from "@reduxjs/toolkit";
import { jwtDecode } from "jwt-decode";
import store, { type AppDispatch, type StoreState } from "..";
import apollo from "../../apollo";
import type {
	GetTokensFromAuth0Query,
	RefreshTokensMutation,
	Tokens,
	User,
} from "../../graphql/generated/graphql";
import GetTokensFromAuth0 from "../../graphql/query/GetTokensFromAuth0";
import RefreshTokens from "../../graphql/query/RefreshTokens";

type AuthState = {
	user: User | null;
	auth0Token: string | null;
	tokens: Tokens | null;
};

const initialState: AuthState = {
	user: null,
	auth0Token: null,
	tokens: null,
};

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
		signOut: (state) => {
			state.user = null;
			state.tokens = null;
		},
	},
	extraReducers: (builder) => {
		builder.addCase(getTokensFromAuth0.pending, (state, action) => {
			state.auth0Token = action.meta.arg.token;
		});
	},
});

export default auth;

export function selectTokens(state: StoreState) {
	return state.auth.tokens;
}

export const selectAccessToken = createSelector(
	selectTokens,
	(tokens) => tokens?.access,
);

export const selectRefreshToken = createSelector(
	selectTokens,
	(tokens) => tokens?.refresh,
);

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
			"failed to refresh tokens: " + errors.map((e) => e.message).join(", "),
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

export const getTokensFromAuth0 = createAsyncThunk<
	GetTokensFromAuth0Query["getTokensFromAuth0"],
	{ token: string },
	{ dispatch: AppDispatch }
>("auth/getTokensFromAuth0", async ({ token }, { dispatch }) => {
	// query for tokens
	const { data, errors } = await apollo.query({
		query: GetTokensFromAuth0,
		variables: { token },
	});
	if (errors) {
		const error = new Error(
			`failed to get tokens from Auth0: ${errors.map((e) => e.message).join(", ")}`,
		);
		console.error(error);
		throw error;
	}

	if (!data?.getTokensFromAuth0)
		throw new Error("failed to get tokens from Auth0");

	// store the tokens in the store
	dispatch(auth.actions.gotTokens(data.getTokensFromAuth0));

	// return tokens
	return data.getTokensFromAuth0;
});
