<template>
  <div class="container">
    <h2 class="title is-3 has-text-centered">Account Details</h2>
  </div>
</template>

<script>
export default {
  name: "Details",
  data: () => {
    return {
      name: "",
      url: "",
      loadingUser: "",
      errors: {
        password: false,
        confirmPassword: false
      }
    };
  },
  async mounted() {
    // fetch user details here
  }
};
</script>

<style lang="scss" scoped></style>
