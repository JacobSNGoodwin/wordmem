<template>
  <div>
    <figure class="image is-128x128">
      <div class="placeholder">
        <img v-if="imageUrl" :src="imageUrl" />
      </div>
    </figure>
    <div class="buttons is-centered my-5">
      <div class="button is-link" @click="openImageSelector">Update Image</div>
    </div>
    <ImageSelector
      :isActive="imageSelectorActive"
      @close="closeImageSelector"
    />
  </div>
</template>

<script>
import ImageSelector from "./ImageSelector";

export default {
  name: "UpdateForm",
  components: {
    ImageSelector
  },
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
      imageFile: null,
      imageSelectorActive: false
    };
  },
  methods: {
    openImageSelector() {
      this.imageSelectorActive = true;
    },
    closeImageSelector() {
      this.imageSelectorActive = false;
    }
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
