import { createRouter, createWebHistory } from "vue-router"
import LoginUser from "./components/User/LoginUser.vue"
import HomePage from "./components/HomePage.vue"
import RegisterUser from "./components/User/RegisterUser.vue"
import ExpensesList from "./components/Expense/ExpensesList.vue"
import SignUp from './components/User/SignUp.vue'
import AuthWrapper from './components/User/AuthWrapper.vue'


const routes = [
    { path: '/', redirect: '/auth' },
    { path: '/auth', component: AuthWrapper },
    { path: '/login', component: LoginUser },
    { path: '/signup', component: SignUp },
    { path: '/register', component: RegisterUser },
    { path: '/home', component: HomePage },
    { path: '/expenses', component: ExpensesList },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const publicPages = ['/login', '/signup'];
    const authRequired = !publicPages.includes(to.path);
    const loggedIn = localStorage.getItem('user');

    if (authRequired && !loggedIn) {
        return next('/login');
    }

    next();
});


export default router;