import "./initComposition"; // this is because of webpack config updated in full VUE3
import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import "./scss/index.scss";
import { provideAuth } from "./store/auth";

Vue.config.productionTip = false;

new Vue({
  router,
  setup() {
    provideAuth();
  },
  render: h => h(App)
}).$mount("#app");
