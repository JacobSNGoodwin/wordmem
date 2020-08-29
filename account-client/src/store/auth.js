import { reactive, provide, inject, toRefs } from "@vue/composition-api";
import jwt_decode from "jwt-decode";
import axios from "axios";

const state = reactive({
  currentUser: null,
  idToken: null,
  refreshToken: null,
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

// in vue3 (as opposed to plugin), we can use the "readonly"
const authStore = {
  state: toRefs(state), // consuming component can destructure withou losing reactivity!
  signin,
  signup
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

// util functions

// authenticate implements common code between signin and signup
const authenticate = async (email, password, endpoint) => {
  state.isLoading = true;

  try {
    const res = await axios.post(endpoint, { email, password });

    const { tokens } = res.data;

    state.idToken = tokens.idToken;
    state.refreshToken = tokens.refreshToken;

    const tokenClaims = jwt_decode(tokens.idToken);

    // set tokens to local storage with expiry (separate function)
    state.currentUser = tokenClaims.user;
  } catch (e) {
    // e.response for non 200
    // e.request for requet errors
    // else some other error
    console.log(e.response);
    state.currentUser = null;
    state.idToken = null;
    state.refreshToken = null;
    state.error = e;
  } finally {
    state.isLoading = false;
  }
};
