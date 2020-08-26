import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home";
import Auth from "../components/Auth";
import Details from "../views/Details";

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
    component: Auth
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
