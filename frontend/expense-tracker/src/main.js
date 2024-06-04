import { createApp } from 'vue'
import App from './App.vue'
import router from './routes.js'
import store from './store.js'
// import axios from 'axios'
const app = createApp(App);

// router.beforeEach(async (to, from, next) => {
//     if (store.getters.isLogged) {
//         try {
//             const response = await axios.get('http://localhost:8090/check-auth')
//             if (response.data.authenticated) {
//                 next()
//             } else {
//                 store.commit('set_login_status', { success: false, message: '' })
//                 next('/login')
//             }
//         } catch (error) {
//             store.commit('set_login_status', { success: false, message: '' })
//             next('/login')
//         }
//     } else {
//         next()
//     }
// })

app.use(router);
app.use(store)
app.mount('#app')
