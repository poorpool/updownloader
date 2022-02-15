import { createRouter, createWebHistory} from "vue-router";
//eslint-disable-next-line no-unused-vars
import Admin from "@/views/Admin";

const home = () => import("../views/Home")

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
  {
    path: "/admin",
    name: "admin",
    component: Admin
  },

]
export const router = createRouter({
  history: createWebHistory(process.env.NODE_ENV === "production" ? "/updown/" : "/"),
  routes: routes
})
