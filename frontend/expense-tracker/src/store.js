
import { createStore } from 'vuex'
import axios from 'axios'
import router from './routes'
import createPersistedState from 'vuex-persistedstate'



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
        },

    },
    actions: {
        async loginUser({ commit }, { username, password }) {
            try {
                const response = await axios.post("http://localhost:8090/login", { username, password }, { withCredentials: true });

                console.log(response)
                if (response.status == 200) {
                    commit('set_login_status', { success: true, message: "Logged in" });
                    router.push("/home")
                }
            } catch (error) {
                commit('set_login_status', { success: false, message: "User not found" })
                console.error('Error calling login API: ', error);
            }

        },
        async registerUser({ commit }, { username, email, password }) {
            try {
                const response = await axios.post("http://localhost:8090/signup", { username, email, password })

                if (response.status == 200) {
                    commit('set_registered_status', { success: true })
                    router.push("/login")
                }
            } catch (error) {
                commit('set_registered_status', { success: false })
                console.error("Error calling the registration API: ", error)
            }
        },
        async logout({ commit }) {
            try {
                await axios.post("http://localhost:8090/logout")
                commit('set_login_status', { success: false })
                router.push("/login")
                console.log("Logged out")
            } catch (err) {
                console.error("Error logging out: ", err)
            }

        }

    },
    getters: {
        isLogged: (state) => state.isLogged,
        loginMessage: (state) => state.loginMessage,
        isRegistered: (state) => state.isRegistered,
    },

    plugins: [createPersistedState()],
});

export default store;

