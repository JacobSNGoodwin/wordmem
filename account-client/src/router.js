import Vue from "vue";
import VueRouter from "vue-router";
import NotFound from "./views/NotFound";
import Auth from "./views/Auth";
import Details from "./views/Details";
import { authStore } from "./store/auth";

Vue.use(VueRouter);

const routes = [
  {
    path: "/authenticate",
    name: "Auth",
    component: Auth
    // component: () =>
    //   import(/* webpackChunkName: "auth" */ "../components/Auth.vue")
  },
  {
    path: "/",
    name: "Details",
    component: Details,
    beforeEnter: requireAuth
  },
  {
    path: "/*",
    name: "NotFound",
    component: NotFound
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

// can't use arrow funcion. I think becuase
// of how the vue instance is injected into the
// route configuration
function requireAuth(to, from, next) {
  // TODO: Get user before checking
  const { currentUser } = authStore;

  if (currentUser.value) {
    next();
  } else {
    next("/authenticate");
  }
}

export default router;
