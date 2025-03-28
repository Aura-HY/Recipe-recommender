import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "",
      redirect: "/login",
    },
    {
      path: "/login",
      name: "login",
      component: () => import("../pages/login.vue"),
      meta: {
        login: true,
      },
    },
    {
      path: "/home",
      name: "home",
      component: () => import("../pages/home.vue"),
      meta: {
        home: true,
      },
    },
  ],
});

export default router;
