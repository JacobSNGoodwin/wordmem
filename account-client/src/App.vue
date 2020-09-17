<template>
  <div id="app">
    <section class="section">
      <Loader v-if="loading" class="my-5" />
      <router-view v-else></router-view>
    </section>
  </div>
</template>

<script>
// import { watchEffect } from "@vue/composition-api";
import { onMounted } from "@vue/composition-api";
import { useAuth } from "./store/auth";
import Loader from "./components/ui/Loader";
export default {
  name: "App",
  components: {
    Loader
  },
  setup() {
    const { currentUser, getUser, loading } = useAuth();

    onMounted(() => {
      getUser();
    });

    // redundant navigation
    // watchEffect(() => {
    //   if (!currentUser.value) {
    //     ctx.root.$router.push("/authenticate");
    //   }
    // });

    return {
      currentUser,
      loading
    };
  }
};
</script>

<style lang="scss" scoped></style>
