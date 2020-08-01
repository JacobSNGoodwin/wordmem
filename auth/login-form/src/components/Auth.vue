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
        <Login :isLogin="isLogin" class="mt-4 mb-4" />
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import Login from "../components/Login";

export default {
  name: "Auth",
  components: { Login },
  setup() {
    const isLogin = ref(true);

    const setIsLogin = (newVal) => {
      isLogin.value = newVal;
    };

    return {
      isLogin,
      setIsLogin,
    };
  },
};
</script>

<style lang="scss" scoped>
.container {
  max-width: 720px;
}
</style>
