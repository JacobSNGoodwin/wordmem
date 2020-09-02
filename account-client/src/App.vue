<template>
  <div id="app">
    <section class="section">
      <router-view></router-view>
    </section>
  </div>
</template>

<script>
import { useAuth } from "./store/auth";
import { onMounted } from "@vue/composition-api";
export default {
  name: "App",
  setup() {
    const { getUser, currentUser, idToken } = useAuth();

    onMounted(async () => {
      await getUser();
    });

    return {
      currentUser,
      idToken
    };
  }
};
</script>

<style lang="scss" scoped></style>
