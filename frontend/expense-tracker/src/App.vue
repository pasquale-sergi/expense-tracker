<template>
  <nav>
    <p v-if="isLogged">You are logged</p>
    <p v-else>You are not logged</p>
    <button v-if="isLogged" @click="logoutUser">Logout</button>
  </nav>

  <login-user v-if="!isLogged"></login-user>

  <home-page v-if="isLogged"></home-page>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import LoginUser from "./components/LoginUser.vue";
// import RegisterUser from './components/RegisterUser.vue';
import HomePage from "./components/HomePage.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    LoginUser,
    HomePage,
  },
  computed: {
    ...mapGetters(["isLogged"]),
  },
  methods: {
    ...mapActions(["logout"]),
    async logoutUser() {
      this.logout();
      try {
        await axios.post("http://localhost:8090/logout");
        console.log("post request to log out sent");
      } catch (err) {
        console.log("Issues with the log out");
      }
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
