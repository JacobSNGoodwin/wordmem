<template>
  <div>
    <figure class="image is-128x128">
      <div class="placeholder">
        <img v-if="imageUrl" :src="imageUrl" />
      </div>
    </figure>
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
  },
  data() {
    return {
      name: this.user.name,
      email: this.user.email,
      website: this.user.website,
      imageUrl: this.user.imageUrl,
      imageFile: null
    };
  }
};
</script>

<style scoped lang="scss">
.image {
  margin: auto;
}

.placeholder {
  width: 128px;
  height: 128px;
  background-color: #f5f5f5;
  border-radius: 64px;

  img {
    height: 128px;
    border-radius: 64px;
  }
}

button {
  width: 120px;
}
</style>
