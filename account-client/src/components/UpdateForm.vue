<template>
  <div class="my-6">
    <figure class="image is-128x128">
      <div class="placeholder">
        <img v-if="imageUrl" :src="imageUrl" />
      </div>
    </figure>
    <div class="buttons is-centered my-6">
      <div class="button is-link" @click="openImageSelector">Update Image</div>
      <div
        @click="deleteUserImage"
        class="button is-danger"
        :class="{ 'is-loading': isDeleteing }"
      >
        Delete Image
      </div>
    </div>
    <div v-if="deleteError">{{ deleteError }}</div>
    <ImageSelector
      :isActive="imageSelectorActive"
      @close="closeImageSelector"
    />
    <!-- Insert rest of form here -->
  </div>
</template>

<script>
import ImageSelector from "./ImageSelector";
import useRequest from "../composables/useRequest";
import { useAuth } from "../store/auth";

export default {
  name: "UpdateForm",
  components: {
    ImageSelector
  },
  props: {
    user: {
      type: Object,
      default: null
    }
  },
  setup() {
    const { idToken } = useAuth();

    const { exec, error: deleteError, loading: isDeleteing } = useRequest({
      url: "/api/image",
      method: "delete",
      headers: {
        Authorization: `Bearer ${idToken.value}`
      }
    });
    return { exec, deleteError, isDeleteing };
  },
  data() {
    return {
      name: this.user.name,
      email: this.user.email,
      website: this.user.website,
      imageUrl: this.user.imageUrl,
      imageSelectorActive: false
    };
  },
  methods: {
    openImageSelector() {
      this.imageSelectorActive = true;
    },
    closeImageSelector() {
      this.imageSelectorActive = false;
    },
    async deleteUserImage() {
      // probably this should be built into the composable
      // as we might have to repeat this callback in other components
      await this.exec();

      if (!this.deleteError) {
        this.imageUrl = "";
      }
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
