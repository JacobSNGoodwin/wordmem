<template>
  <div>
    <h2 class="title is-3 has-text-centered">
      {{ isLogin ? "Login" : "Sign Up" }}
    </h2>
    <form action="" method="post" @submit="validateForm">
      <div class="field my-5">
        <div class="control">
          <input
            class="input is-rounded has-text-weight-bold is-centered"
            type="email"
            v-model="email"
            placeholder="Email Address"
          />
        </div>
        <p class="help is-danger has-text-centered">
          The entered email is invalid
        </p>
      </div>
      <div class="field my-5">
        <div class="control">
          <input
            class="input is-rounded has-text-weight-bold"
            type="password"
            v-model="password"
            placeholder="Password"
          />
          <p class="help is-danger has-text-centered">
            Password must be between 6 and 30 characters
          </p>
        </div>
      </div>
      <div v-if="!isLogin" class="field my-5">
        <div class="control">
          <input
            class="input is-rounded has-text-weight-bold"
            type="password"
            v-model="confirmPassword"
            placeholder="Confirm Password"
          />
        </div>
        <p class="help is-danger has-text-centered">
          Passwords do not match
        </p>
      </div>
      <div class="buttons is-centered mt-6">
        <button type="submit" class="button is-info is-rounded">
          {{ isLogin ? "Login" : "Sign Up" }}
        </button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  name: "Login",
  components: {},
  props: {
    isLogin: Boolean,
  },
  data: () => {
    return {
      email: "",
      password: "",
      confirmPassword: "",
      errors: {
        email: null,
        password: null,
        confirmPassword: null,
      },
    };
  },
  methods: {
    validateForm: () => {},
  },
};
</script>

<style scoped lang="scss">
.field {
  max-width: 480px;
  margin-left: auto;
  margin-right: auto;
}
</style>
