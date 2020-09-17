<template>
  <div class="container">
    <div class="card">
      <div class="card-content">
        <div class="tabs is-centered is-medium">
          <ul>
            <li @click="setIsLogin(true)" :class="{ 'is-active': isLogin }">
              <a>Login</a>
            </li>
            <li @click="setIsLogin(false)" :class="{ 'is-active': !isLogin }">
              <a>Sign Up</a>
            </li>
          </ul>
        </div>
        <h1 class="has-text-centered title is-2 mb-6">
          Not Too Evil, Inc.
        </h1>
        <LoginForm
          :isLogin="isLogin"
          :isFetchingData="isLoading"
          class="mt-4 mb-4"
          @authSubmitted="authSubmitted"
        />
        <div v-if="error" class="has-text-centered">
          <p class="has-text-danger">{{ error.message }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, watchEffect } from "@vue/composition-api";
import LoginForm from "../components/LoginForm";
import { useAuth } from "../store/auth";

export default {
  name: "Auth",
  components: { LoginForm },
  setup(_, ctx) {
    const isLogin = ref(true);

    const setIsLogin = newVal => {
      isLogin.value = newVal;
    };

    const { currentUser, error, isLoading, signin, signup } = useAuth();

    const authSubmitted = ({ email, password }) => {
      isLogin.value ? signin(email, password) : signup(email, password);
    };

    watchEffect(() => {
      if (currentUser.value) {
        ctx.root.$router.push({ name: "Details" });
      }
    });

    return {
      isLogin,
      setIsLogin,
      authSubmitted,
      currentUser,
      error,
      isLoading
    };
  }
};
</script>

<style lang="scss" scoped>
.container {
  max-width: 720px;
}
</style>
