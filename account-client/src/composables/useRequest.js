import axios from "axios";
import { ref } from "@vue/composition-api";

// a generic request function.
// it will also help us to differentiate
// between request/response errors
const useRequest = (url, method, body) => {
  const errors = ref([]);
  const data = ref(null);
  const loading = ref(false);

  const exec = async () => {
    try {
      loading.value = true;
      errors.value = [];
      data.value = null;

      const response = await axios[method](url, body);
      data.value = response.data;
    } catch (error) {
      console.log("error: ", error);
    } finally {
      loading.value = false;
    }
  };

  return {
    exec,
    errors,
    data,
    loading
  };
};

export default useRequest;
