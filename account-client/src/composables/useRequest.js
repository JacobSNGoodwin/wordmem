import { ref } from "@vue/composition-api";
import { doRequest } from "../util";

// request function to wrap the doRequest util method
// this basically allows us to also add some state
const useRequest = reqOptions => {
  const error = ref(null);
  const data = ref(null);
  const loading = ref(false);

  const exec = async () => {
    loading.value = true;

    const resp = await doRequest(reqOptions);
    data.value = resp.value;
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
