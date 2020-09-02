import axios from "axios";
import jwt_decode from "jwt-decode";

// doRequest is a helper function for
// handling axios responses
export const doRequest = async (url, method, body) => {
  let error;
  let data;

  try {
    const response = await axios[method](url, body);
    data = response.data;
  } catch (e) {
    if (e.response) {
      error = e.response.data.error;
    } else if (e.request) {
      error = e.request;
    } else {
      error = e;
    }
  }

  return {
    data,
    error
  };
};

// storeTokens utility for storing idAndRefreshToken
export const storeTokens = (idToken, refreshToken) => {
  localStorage.setItem("__evilCorpId", idToken);
  localStorage.setItem("__evilCorpRf", refreshToken);
};

// gets the token's payload, and returns null
// if invalid
export const getTokenPayload = token => {
  if (!token) {
    return null;
  }

  const tokenClaims = jwt_decode(token);

  if (tokenClaims.exp >= Date.now()) {
    return null;
  }

  return tokenClaims;
};