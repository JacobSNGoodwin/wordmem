import "./initComposition"; // this is because of webpack config updated in full VUE3
import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import "./scss/index.scss";

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
