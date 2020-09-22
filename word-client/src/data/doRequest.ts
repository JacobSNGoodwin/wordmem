import axios, { AxiosRequestConfig } from "axios";

// handling axios responses - reqOptions follow axios req config
export const doRequest = async <T>(reqOptions: AxiosRequestConfig) => {
  let error: Error | undefined;
  let data: T | undefined;

  console.log("In do request...");

  try {
    const response = await axios.request<T>(reqOptions);
    data = response.data;

    console.log("Successfull response data");
  } catch (e) {
    console.log("Error data");
    if (e.response) {
      error = e.response.data.errors;
    } else if (e.request) {
      error = e.request;
    } else {
      error = e;
    }
  }

  console.log(data);
  console.log(error);

  return {
    data,
    error,
  };
};
