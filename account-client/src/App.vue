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
    const { getUser, currentUser } = useAuth();

    onMounted(async () => {
      await getUser();
    });

    return {
      currentUser
    };
  }
};
</script>

<style lang="scss" scoped></style>
