import { ref } from "@vue/composition-api";
import { doRequest } from "../util";

// request function to wrap the doRequest util method
// this basically allows us to also add some state
const useRequest = reqOptions => {
  const error = ref(null);
  const data = ref(null);
  const loading = ref(false);

  // optional data param to merge into request options
  const exec = async reqData => {
    loading.value = true;
    error.value = null;

    if (data) {
      reqOptions = {
        data: reqData,
        ...reqOptions
      };
    }

    const resp = await doRequest(reqOptions);

    data.value = resp.data;
    error.value = resp.error;
    loading.value = false;
  };

  return {
    exec,
    error,
    data,
    loading
  };
};

export default useRequest;
