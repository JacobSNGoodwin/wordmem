import create, { SetState } from "zustand";
import {
  doRequest,
  getTokenPayload,
  IdTokenClaims,
  RefreshTokenClaims,
  storeTokens,
} from "../util";

export type User = {
  uid: string;
  email: string;
  name: string;
  imageUrl: string;
  website: string;
};

// make properties readonly
type AuthState = {
  currentUser?: User;
  idToken?: string;
  isLoading: boolean;
  error?: Error;
  getUser: (forceRefresh: boolean) => Promise<void>;
};

export const useAuth = create<AuthState>((set) => {
  return {
    currentUser: undefined,
    idToken: "",
    isLoading: false,
    error: undefined,
    getUser: (forceRefresh: boolean = false) => getUser({ set, forceRefresh }),
  };
});

const getUser = async (options: {
  set: SetState<AuthState>;
  forceRefresh: boolean;
}) => {
  const { set, forceRefresh } = options;
  set({
    isLoading: true,
    error: undefined,
  });

  // ghetto ass way to convert possible null to undefined since
  // I likes working with undefineds onlyz
  const idToken = localStorage.getItem("__evilCorpId") ?? undefined; // add env variable globally
  const idTokenClaims = getTokenPayload<IdTokenClaims>(idToken);

  // if we have a valid idToken, set the user (use spread with merged obj?)
  if (idTokenClaims && !forceRefresh) {
    set({
      idToken: idToken,
      currentUser: idTokenClaims.user,
      isLoading: false,
    });

    return;
  }

  // we don't have a valid or non-expired idToken, or we want to force refresh
  // so we try the refresh token
  const refreshToken = localStorage.getItem("__evilCorpRf") ?? undefined;
  const refreshTokenClaims = getTokenPayload<RefreshTokenClaims>(refreshToken);

  // return setting user to null if no refresh token
  if (!refreshTokenClaims) {
    set({
      currentUser: undefined,
      idToken: undefined,
      isLoading: false,
    });

    return;
  }

  interface TokenResponse {
    tokens: {
      idToken: string;
      refreshToken: string;
    };
  }

  const { data, error } = await doRequest<TokenResponse>({
    url: "/api/account/tokens",
    method: "post",
    data: {
      refreshToken,
    },
  });

  // failure to get a response from the endpoint
  if (error || !data) {
    set({
      currentUser: undefined,
      idToken: undefined,
      error: error || Error("Could not fetch tokens"),
      isLoading: false,
    });

    return;
  }

  const { tokens } = data;
  storeTokens(tokens.idToken, tokens.refreshToken);
  const tokenClaims = getTokenPayload<IdTokenClaims>(tokens.idToken);

  set({
    idToken: tokens.idToken,
    currentUser: tokenClaims!.user, // we'll just be generous an assume this is a valid token
    isLoading: false,
  });
};
