<template>
  <div>
    {{ user }}
  </div>
</template>

<script>
export default {
  name: "UpdateForm",
  components: {},
  props: {
    user: {
      type: Object,
      default: null
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
