<template>
  <Auth />
</template>

<script>
// need to load auth state here and nav to details if we're authenticated
// logout button and profile photo in navbar if logged in
import Auth from "../components/Auth";

export default {
  name: "Home",
  components: {
    Auth
  }
};
</script>

<style lang="scss" scoped></style>
