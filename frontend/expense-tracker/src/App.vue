<template>
  <nav class="navbar">
    <p class="user-welcome" v-if="isLogged">
      Welcome Back {{ capitalizeFirstLetter(username) }}
    </p>

    <button v-if="isLogged" @click="logoutUser">Logout</button>
  </nav>

  <login-user v-if="!isLogged"></login-user>

  <home-page v-if="isLogged"></home-page>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import LoginUser from "./components/User/LoginUser.vue";
// import RegisterUser from './components/RegisterUser.vue';
import HomePage from "./components/HomePage.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    LoginUser,
    HomePage,
  },
  data() {
    return {
      username: "",
    };
  },
  created() {
    this.getUsername();
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
    async getUsername() {
      try {
        const response = await axios.get("http://localhost:8090/username");
        this.username = response.data.username;
      } catch (err) {
        console.log("Error retrieving the username");
      }
    },
    capitalizeFirstLetter(str) {
      return str.charAt(0).toUpperCase() + str.slice(1);
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
.navbar {
  display: flex;
  border: solid 2px solid;
  padding: 20px;
}

.user-welcome {
  font-size: 30px;
  margin-left: 45%;
}
</style>
