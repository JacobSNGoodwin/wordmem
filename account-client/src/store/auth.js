import { reactive, provide, inject, toRefs } from "@vue/composition-api";

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
  const req = new Request(endpoint, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      email,
      password
    })
  });

  state.isLoading = true;

  try {
    const res = await fetch(req);
    const { user, tokens } = await res.json();

    state.idToken = tokens.idToken;
    state.refreshToken = tokens.refreshToken;
    state.currentUser = user;
  } catch (e) {
    console.log(e);
    state.currentUser = null;
    state.idToken = null;
    state.refreshToken = null;
    state.error = new Error("Error fetching user");
  } finally {
    state.isLoading = false;
  }
};
