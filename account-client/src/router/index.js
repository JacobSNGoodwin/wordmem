import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home";
import Details from "../views/Details";
// import Auth from "../components/Auth";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/authenticate",
    name: "Auth",
    component: () =>
      import(/* webpackChunkName: "auth" */ "../components/Auth.vue")
  },
  {
    path: "/details",
    name: "Details",
    component: Details
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
