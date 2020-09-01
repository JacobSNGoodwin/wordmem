<template>
  <div>
    <Details />
  </div>
</template>

<script>
import Details from "../components/Details";
import { useAuth } from "../store/auth";
import { onMounted } from "@vue/composition-api";

export default {
  name: "Home",
  components: {
    Details
  },
  setup(_, ctx) {
    const { currentUser, getUser } = useAuth();

    onMounted(async () => {
      await getUser();

      if (!currentUser.value) {
        ctx.root.$router.push("/authenticate");
      }
    });

    return {
      currentUser
    };
  }
};
</script>

<style lang="scss" scoped></style>
