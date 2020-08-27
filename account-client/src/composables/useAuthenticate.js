import { useFetch } from "vue-composable";

export const useAuthenticate = () => {
  const { loading, error, json, exec } = useFetch();

  // will return a callable function to execute the login
  const authenticate = authOptions => {
    const { email, password, url } = authOptions;

    // create post request to url
    const req = new Request(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        email,
        password
      })
    });

    return exec(req);
  };

  return {
    loading,
    error,
    json,
    authenticate
  };
};
