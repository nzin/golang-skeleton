import { createWebHashHistory, createRouter } from "vue-router";
import TodosApp from "@/components/TodosApp.vue";
import TodoApp from "@/components/TodoApp.vue";

const routes = [
  {
    path: "/",
    name: "todos",
    component: TodosApp,
  },
  {
    path: "/todos/:todoid",
    name: "todo",
    component: TodoApp,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;