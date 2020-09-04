<template>
  <div class="file has-name is-fullwidth">
    <label class="file-label">
      <input
        class="file-input"
        type="file"
        accept=".jpg,.jpeg,.png"
        @change="fileChanged"
      />
      <span class="file-cta">
        <span class="file-label">
          Choose a file…
        </span>
      </span>
      <span class="file-name">
        {{ selectedFile && selectedFile.name }}
      </span>
    </label>
  </div>
</template>

<script>
export default {
  name: "FileSelector",
  data() {
    return {
      selectedFile: null
    };
  },
  methods: {
    fileChanged(e) {
      const files = e.target.files || e.dataTransfer.files;
      if (!files.length) {
        // no file
        this.selectedFile = null;
        this.$emit("fileChanged", null);
        return;
      }
      const selectedFile = files[0];
      this.selectedFile = selectedFile;
      this.$emit("fileChanged", selectedFile);
    }
  }
};
</script>
