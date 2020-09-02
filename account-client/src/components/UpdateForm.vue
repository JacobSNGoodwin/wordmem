<template>
  <div>
    <h2 class="title is-3 has-text-centered">
      Update User Form
    </h2>
  </div>
</template>

<script>
export default {
  name: "UpdateForm",
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
