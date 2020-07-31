<template>
  <h2 class="title is-3 has-text-centered">{{ isLogin ? "Login" : "Sign Up" }}</h2>
</template>

<script>
export default {
  name: "Login",
  props: {
    isLogin: Boolean,
  },
};
</script>
