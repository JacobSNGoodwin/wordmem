<template>
  <div id="app">
    <section class="section">
      <router-view></router-view>
    </section>
  </div>
</template>

<script>
import { provideAuth } from "./store/auth";
export default {
  name: "App",
  setup() {
    provideAuth();
  }
};
</script>

<style lang="scss" scoped></style>
