<template>
  <div>
    <Auth />
  </div>
</template>

<script>
// need to load auth state here and nav to details if we're authenticated
// logout button and profile photo in navbar if logged in
import Auth from "../components/Auth";
import { onMounted } from "@vue/composition-api";
import { useAuth } from "../store/auth";

export default {
  name: "Home",
  components: {
    Auth
  },
  setup() {
    const { state, initUser } = useAuth();

    onMounted(async () => {
      console.log("Home Mounted in setup, fools");
      console.log("Before initUser(): ", state);

      await initUser();

      console.log("After initUser(): ", state);
    });
  }
};
</script>

<style lang="scss" scoped></style>
