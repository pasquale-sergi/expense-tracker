
import { createStore } from 'vuex'
import axios from 'axios'
import router from './routes'


const store = createStore({
    state: {
        isLogged: false,
        isRegistered: false,
        loginMessage: '',

    },
    mutations: {
        set_login_status(state, { success, message }) {
            state.isLogged = success;
            state.loginMessage = message;
        },
        set_registered_status(state, success) {
            state.isRegistered = success;
        }
    },
    actions: {
        async loginUser({ commit }, { username, password }) {
            try {
                const response = await axios.post("http://localhost:8090/login", { username, password });
                commit('set_login_status', { success: response.data.success, message: response.data.message })
                console.log(response)
                if (response.data.success) {
                    router.push("/home")
                }


            } catch (error) {
                commit('set_login_status', { success: false, message: "User not found" })
                console.error('Error calling login API: ', error);
            }

        },
        async registerUser({ commit }, { username, email, password }) {
            try {
                const response = await axios.post("http://localhost:8090/register", { username, email, password })
                commit('set_registered_status', { success: response.data.success })
                if (response.data.success) {
                    router.push("/home")
                }
            } catch (error) {
                commit('set_registered_status', { success: false })
                console.error("Error calling the registration API: ", error)
            }
        }

    },
    getters: {
        isLogged: (state) => state.isLogged,
        loginMessage: (state) => state.loginMessage,
        isRegistered: (state) => state.isRegistered,
    }
});

export default store;

