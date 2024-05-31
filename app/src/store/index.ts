import { configureStore } from "@reduxjs/toolkit";
import auth from "./slices/auth";

const store = configureStore({ reducer: { auth: auth.reducer } });

export type StoreState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export default store;
