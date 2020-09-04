<template>
  <div class="modal" :class="{ 'is-active': isActive }">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Update Profile Image</p>
      </header>
      <section class="modal-card-body">
        <p>Some stuffs</p>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success">Save</button>
        <button class="button" @click="close">Cancel</button>
      </footer>
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageSelector",
  props: {
    isActive: {
      type: Boolean,
      required: true
    }
  },
  methods: {
    close() {
      this.$emit("close");
    }
  }
};
</script>

<style scoped></style>
