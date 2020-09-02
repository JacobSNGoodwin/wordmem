<template>
  <div>
    <h2 class="title is-3 has-text-centered">
      {{ isLogin ? "Login" : "Sign Up" }}
    </h2>
    <ValidationObserver v-slot="{ handleSubmit, invalid }">
      <form novalidate="true" @submit.prevent="handleSubmit(submitForm)">
        <div class="field my-5">
          <div class="control">
            <ValidationProvider name="email" rules="required" v-slot="v">
              <input
                class="input is-rounded has-text-weight-bold is-centered"
                type="email"
                v-model="email"
                placeholder="Email Address"
              />
              <div
                v-if="v.touched && v.invalid"
                class="help is-danger has-text-centered"
              >
                <p v-for="error in v.errors" :key="error">
                  {{ error }}
                </p>
              </div>
            </ValidationProvider>
          </div>
        </div>
        <div class="field my-5">
          <div class="control">
            <ValidationProvider
              name="password"
              rules="required|min:6|max:30"
              v-slot="v"
            >
              <input
                class="input is-rounded has-text-weight-bold"
                type="password"
                vid="password"
                v-model="password"
                placeholder="Password"
              />
              <div
                v-if="v.touched && v.invalid"
                class="help is-danger has-text-centered"
              >
                <p v-for="error in v.errors" :key="error">
                  {{ error }}
                </p>
              </div>
            </ValidationProvider>
          </div>
        </div>
        <div v-if="!isLogin" class="field my-5">
          <div class="control">
            <ValidationProvider
              name="confirmPassword"
              rules="required|confirmed:password"
              v-slot="v"
            >
              <input
                class="input is-rounded has-text-weight-bold"
                type="password"
                v-model="confirmPassword"
                placeholder="Confirm Password"
              />
              <div
                v-if="v.touched && v.invalid"
                class="help is-danger has-text-centered"
              >
                <p v-for="error in v.errors" :key="error">
                  {{ error }}
                </p>
              </div>
            </ValidationProvider>
          </div>
        </div>
        <div class="buttons is-centered mt-6">
          <button
            type="submit"
            :disabled="invalid"
            class="button is-info is-rounded"
            :class="{ 'is-loading': isFetchingData }"
          >
            {{ isLogin ? "Login" : "Sign Up" }}
          </button>
        </div>
      </form>
    </ValidationObserver>
  </div>
</template>

<script>
export default {
  name: "Login",
  components: {},
  props: {
    isLogin: {
      type: Boolean,
      default: true
    },
    isFetchingData: {
      type: Boolean,
      default: false
    }
  },
  data: () => {
    return {
      email: "",
      password: "",
      confirmPassword: ""
    };
  },
  methods: {
    submitForm() {
      this.$emit("authSubmitted", {
        email: this.email,
        password: this.password
      });
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

button {
  width: 120px;
}
</style>
