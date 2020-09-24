<template>
  <div class="container">
    <h2 class="title is-3 has-text-centered">Account Details</h2>
    <Loader v-if="loading" class="my-6" />
    <UpdateForm v-if="data && !loading" :user="data.user" />
    <div class="buttons is-centered">
      <button
        @click="signout"
        class="button is-rounded is-danger"
        :class="{ 'is-loading': authLoading }"
      >
        Sign Out
      </button>
    </div>
    <p class="has-text-danger has-text-centered">
      Warning: This will sign you out of all devices!
    </p>
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
  setup(_, ctx) {
    const {
      idToken,
      signout: authSignout,
      isLoading: authLoading,
      error: authError
    } = useAuth();

    const { data, error, loading, exec } = useRequest({
      url: "/api/account/me",
      method: "get",
      headers: {
        Authorization: `Bearer ${idToken.value}`
      }
    });

    exec();

    const signout = () => {
      authSignout().then(() => {
        ctx.root.$router.push("/authenticate");
      });
    };

    return { data, error, loading, signout, authLoading, authError, idToken };
  }
};
</script>

<style lang="scss" scoped></style>
