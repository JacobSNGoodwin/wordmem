import { reactive, provide, inject, toRefs } from "@vue/composition-api";
import jwt_decode from "jwt-decode";
import { storeTokens, getTokenPayload, doRequest } from "../util";

const state = reactive({
  currentUser: null,
  isLoading: false,
  error: null
});

// signin/signup reach out to api endpoint
// on successful request, it sets the current use
// along with the id and refresh tokens
const signin = async (email, password) =>
  await authenticate(email, password, "/api/signin");

const signup = async (email, password) =>
  await authenticate(email, password, "/api/signup");

// this method can be used on page load
// to check for a current valid user idToken (short-lived)
// If there is no short lived token, it checks for a long-lived token
// and submits this to the refresh token endpoint
const getUser = async () => {
  state.isLoading = true;
  state.error = null;

  // check for idToken
  const idToken = localStorage.getItem("__evilCorpId");
  const idTokenClaims = getTokenPayload(idToken);

  // if we have a valid idToken, set the user
  if (idTokenClaims) {
    state.currentUser = idTokenClaims.user;
    state.isLoading = false;
    return;
  }

  // we don't have a valid or non-expired idToken
  // so we try the refresh token
  const refreshToken = localStorage.getItem("__evilCorpRf");
  const refreshTokenClaims = getTokenPayload(refreshToken);

  // return setting user to null if no refresh token
  if (!refreshTokenClaims) {
    state.currentUser = null;
    state.isLoading = false;

    return;
  }

  // try refresh endpoint

  const { data, error } = await doRequest("/api/tokens", "post", {
    refreshToken
  });

  // failure to get a response from the endpoint
  if (error) {
    state.currentUser = null;
    state.error = error;
    state.isLoading = false;
    return;
  }

  const { tokens } = data;
  storeTokens(tokens.idToken, tokens.refreshToken);
  const tokenClaims = jwt_decode(tokens.idToken);

  // set tokens to local storage with expiry (separate function)
  state.currentUser = tokenClaims.user;
  state.isLoading = false;
};

// const signout = async(idToken) =>

// in vue3 (as opposed to plugin), we can use the "readonly"/ Mp readonly option in preview
export const authStore = {
  ...toRefs(state), // consuming component can destructure withou losing reactivity!
  signin,
  signup,
  getUser
};

// Create functions so the store can be
// injected down the application tree
const StoreSymbol = Symbol("authStore");

export function provideAuth() {
  provide(StoreSymbol, authStore);
}

export function useAuth() {
  const store = inject(StoreSymbol);

  if (!store) {
    throw new Error("Auth store have not been instantiated!");
  }

  return store;
}

// authenticate implements common code between signin and signup
const authenticate = async (email, password, url) => {
  state.isLoading = true;
  state.error = null;

  const { data, error } = await doRequest(url, "post", { email, password });

  if (error) {
    state.error = error;
    state.isLoading = false;
    return;
  }

  const { tokens } = data;

  storeTokens(tokens.idToken, tokens.refreshToken);

  const tokenClaims = jwt_decode(tokens.idToken);

  // set tokens to local storage with expiry (separate function)
  state.currentUser = tokenClaims.user;
  state.isLoading = false;
};
