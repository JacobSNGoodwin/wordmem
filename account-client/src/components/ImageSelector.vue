<template>
  <div class="modal" :class="{ 'is-active': isActive }">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Update Profile Image</p>
      </header>
      <section class="modal-card-body">
        <div>
          <VueCroppie
            ref="croppieRef"
            :boundary="{ width: 258, height: 258 }"
            :viewport="{ width: 256, height: 256, type: 'circle' }"
            :enableResize="false"
            mouseWheelZoom="ctrl"
            :showZoomer="true"
          />
        </div>
        <FileSelector @fileChanged="fileChanged" />
      </section>
      <footer class="modal-card-foot">
        <button
          class="button is-link"
          :class="{ 'is-loading': isUploading }"
          @click="cropAndUpload"
        >
          Upload
        </button>
        <button class="button" @click="close">Cancel</button>
      </footer>
    </div>
  </div>
</template>

<script>
import { watchEffect } from "@vue/composition-api";
import FileSelector from "./ui/FileSelector";
import useRequest from "../composables/useRequest";
import { useAuth } from "../store/auth";

export default {
  name: "ImageSelector",
  components: {
    FileSelector
  },
  props: {
    isActive: {
      type: Boolean,
      required: true
    }
  },
  setup(_, { emit }) {
    const { idToken } = useAuth();

    const {
      exec: uploadImage,
      data: uploadData,
      error: uploadError,
      loading: isUploading
    } = useRequest({
      url: "/api/image",
      method: "post",
      headers: {
        Authorization: `Bearer ${idToken.value}`,
        "Content-Type": "multipart/form-data"
      }
    });

    watchEffect(() => {
      if (uploadData.value) {
        emit("imageUrlUpdated", uploadData.value.imageUrl);
      }
    });

    return { uploadImage, uploadData, uploadError, isUploading };
  },
  methods: {
    close() {
      this.$emit("close");
    },
    fileChanged(file) {
      const reader = new FileReader();

      reader.onload = e => {
        this.$refs.croppieRef.bind({
          url: e.target.result
        });
      };

      reader.readAsDataURL(file);
    },
    cropAndUpload() {
      const cropOptions = {
        type: "blob",
        size: "viewport",
        format: "png"
      };

      this.$refs.croppieRef.result(cropOptions, output => {
        const formData = new FormData();
        formData.append("imageFile", output);
        this.uploadImage(formData);
      });
    }
  }
};
</script>

<style scoped>
.file {
  justify-content: center;
}
.modal-card-foot {
  justify-content: center;
}
</style>