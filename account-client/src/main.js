import "./composition"; // this is because of webpack config updated in full VUE3
import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import "./validator";
import VueCroppie from "vue-croppie";
import "croppie/croppie.css";
import "./scss/index.scss";
import { provideAuth } from "./store/auth";

Vue.config.productionTip = false;
Vue.use(VueCroppie);

new Vue({
  router,
  setup() {
    provideAuth();
  },
  render: h => h(App)
}).$mount("#app");
