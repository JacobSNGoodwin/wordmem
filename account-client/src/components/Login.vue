<template>
  <div>
    <h2 class="title is-3 has-text-centered">
      {{ isLogin ? "Login" : "Sign Up" }}
    </h2>
    <form novalidate="true" @submit="validateForm">
      <div class="field my-5">
        <div class="control">
          <input
            class="input is-rounded has-text-weight-bold is-centered"
            type="email"
            v-model="email"
            placeholder="Email Address"
          />
        </div>
        <p v-if="errors.email" class="help is-danger has-text-centered">
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
          <p v-if="errors.password" class="help is-danger has-text-centered">
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
        <p
          v-if="errors.confirmPassword"
          class="help is-danger has-text-centered"
        >
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
// example of using classic options API alongside vue composition API
import { useFetch } from "vue-composable";
export default {
  name: "Login",
  components: {},
  props: {
    isLogin: Boolean
  },
  setup() {
    const { json: data, loading, exec } = useFetch();

    return {
      data,
      loading,
      exec
    };
  },
  data: () => {
    return {
      email: "",
      password: "",
      confirmPassword: "",
      errors: {
        email: false,
        password: false,
        confirmPassword: false
      }
    };
  },
  methods: {
    async validateForm(event) {
      // prevent submission if there are any errors
      event.preventDefault();

      // clear previous errors
      this.errors.email = false;
      this.errors.password = false;
      this.errors.confirmPassword = false;

      const isEmailValid = this.isEmailValid(this.email);
      const isPasswordValid =
        this.password.length >= 6 && this.password.length <= 30;
      const doPasswordMatch =
        this.password === this.confirmPassword || this.isLogin;

      if (!isEmailValid) {
        this.errors.email = true;
      }

      if (!isPasswordValid) {
        this.errors.password = true;
      }

      if (!doPasswordMatch) {
        this.errors.confirmPassword = true;
      }

      if (isEmailValid && isPasswordValid && doPasswordMatch) {
        const req = new Request("/api/signin", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password
          })
        });
        this.exec(req);
      }
    },
    isEmailValid: email => {
      const emailExp = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

      return emailExp.test(email);
    }
  }
};
</script>

<style scoped lang="scss">
.field {
  max-width: 480px;
  margin-left: auto;
  margin-right: auto;
}
</style>
