import { createRouter, createWebHistory } from "vue-router"
import LoginUser from "./components/LoginUser.vue"
import HomePage from "./components/HomePage.vue"
import RegisterUser from "./components/RegisterUser.vue"

const routes = [
    { path: '/', redirect: '/login' },
    {
        path: '/login',  // Note: It's common to use a simple path like '/login'
        component: LoginUser
    },
    { path: '/register', component: RegisterUser },
    {
        path: '/home',
        component: HomePage
    }

];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;