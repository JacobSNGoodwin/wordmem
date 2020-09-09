<template>
  <div class="container">
    <h2 class="title is-3 has-text-centered">Account Details</h2>
    <Loader v-if="loading" class="my-6" />
    <UpdateForm v-if="data && !loading" :user="data.user" />
  </div>
</template>

<script>
import { useAuth } from "../store/auth";
import useRequest from "../composables/useRequest";
import UpdateForm from "../components/UpdateForm";
import Loader from "../components/ui/Loader";
export default {
  name: "Details",
  components: {
    UpdateForm,
    Loader
  },
  setup() {
    const { idToken } = useAuth();
    const { data, error, loading, exec } = useRequest({
      url: "/api/account/me",
      method: "get",
      headers: {
        Authorization: `Bearer ${idToken.value}`
      }
    });

    exec();

    return { data, error, loading };
  }
};
</script>

<style lang="scss" scoped></style>
