<template>
  <div id="app">
    <section class="section">
      <Auth />
    </section>
  </div>
</template>

<script>
import Auth from "./components/Auth.vue";
export default {
  name: "App",
  components: {
    Auth
  }
};
</script>

<style lang="scss" scoped></style>
