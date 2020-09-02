import Vue from "vue";
import VueCompositionAPI from "@vue/composition-api";

// This is a necessary workaround to be able to use global store
// see - https://stackoverflow.com/questions/61885716/uncaught-error-vue-composition-api-must-call-vue-useplugin-before-using-any
Vue.use(VueCompositionAPI);
