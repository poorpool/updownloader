import { createRouter, createWebHistory} from "vue-router";

const home = () => import("../Home")

const routes = [
  {
    path: "/:code",
    name: "identified",
    component: home
  },
  {
    path: "/",
    name: "home",
    component: home
  },
]
export const router = createRouter({
  history: createWebHistory(),
  routes: routes
})
