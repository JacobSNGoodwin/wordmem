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
        <h1 class="has-text-centered title is-2 mb-6">Not Too Evil, Inc.</h1>
        <Login
          :isLogin="isLogin"
          :isFetchingData="state.isLoading"
          class="mt-4 mb-4"
          @authSubmitted="authSubmitted"
        />
        <div v-if="state.error">
          <p>{{ state.error.message }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "@vue/composition-api";
import Login from "../components/Login";
import { useAuth } from "../store/auth";

export default {
  name: "Auth",
  components: { Login },
  setup() {
    const isLogin = ref(true);

    const setIsLogin = newVal => {
      isLogin.value = newVal;
    };

    const { state, signin, signup } = useAuth();

    const authSubmitted = ({ email, password }) => {
      isLogin.value ? signin(email, password) : signup(email, password);
    };

    return {
      isLogin,
      setIsLogin,
      authSubmitted,
      state
    };
  }
};
</script>

<style lang="scss" scoped>
.container {
  max-width: 720px;
}
</style>
