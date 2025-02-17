import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ReportView from "@/views/ReportView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/book",
      name: "book",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/BookView.vue"),
    },
    {
      path: "/seat",
      name: "seat",
      component: () => import("../views/SeatView.vue"),
    },
    {
      path: "/payment",
      name: "payment",
      component: () => import("../views/PaymentsView.vue"),
    },
    {
      path: "/report/:reference",
      name: "report",
      component: ReportView,
    },
  ],
});

export default router;
