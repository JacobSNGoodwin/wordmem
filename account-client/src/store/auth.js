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
const signin = async (email, password) => {
  const res = await authenticate(email, password, "/api/signin");

  console.log("Signin response: ", await res.json());
};

const signup = async (email, password) => {
  const res = await authenticate(email, password, "/api/signup");

  console.log("Signup response: ", await res.json());
};

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

  const res = await fetch(req);

  state.isLoading = false;

  return res;
};
