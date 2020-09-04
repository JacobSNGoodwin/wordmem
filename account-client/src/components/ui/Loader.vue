<template>
  <div class="spinner" :style="{ width: `${width}px`, height: `${height}px` }">
    <div class="double-bounce1" :style="{ backgroundColor: color }"></div>
    <div class="double-bounce2" :style="{ backgroundColor: color }"></div>
  </div>
</template>

<script>
export default {
  name: "Loader",
  props: {
    width: {
      type: Number,
      default: 80
    },
    height: {
      type: Number,
      default: 80
    },
    color: {
      type: String,
      default: "#3273dc"
    }
  }
};
</script>

<style scoped>
.spinner {
  position: relative;
  margin-left: auto;
  margin-right: auto;
}

.double-bounce1,
.double-bounce2 {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background-color: rgb(4, 0, 255);
  opacity: 0.6;
  position: absolute;
  top: 0;
  left: 0;

  -webkit-animation: sk-bounce 2s infinite ease-in-out;
  animation: sk-bounce 2s infinite ease-in-out;
}

.double-bounce2 {
  -webkit-animation-delay: -1s;
  animation-delay: -1s;
}

@-webkit-keyframes sk-bounce {
  0%,
  100% {
    -webkit-transform: scale(0);
  }
  50% {
    -webkit-transform: scale(1);
  }
}

@keyframes sk-bounce {
  0%,
  100% {
    transform: scale(0);
    -webkit-transform: scale(0);
  }
  50% {
    transform: scale(1);
    -webkit-transform: scale(1);
  }
}
</style>
