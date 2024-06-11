<template>
  <nav class="navbar" v-if="isLogged">
    <div class="navbar-content">
      <p class="user-welcome">
        Welcome Back, {{ capitalizeFirstLetter(username) }}
      </p>
      <button @click="logoutUser" class="logout-btn">Logout</button>
    </div>
  </nav>
  <div class="auth-box" v-if="!isLogged">
    <auth-wrapper @login-success="onLoginSuccess"></auth-wrapper>
  </div>
  <home-page v-if="isLogged"></home-page>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import AuthWrapper from "./components/User/AuthWrapper.vue";
import HomePage from "./components/HomePage.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    HomePage,
    AuthWrapper,
  },
  data() {
    return {
      username: "",
      isLoginView: true,
    };
  },
  computed: {
    ...mapGetters(["isLogged"]),
  },
  created() {
    this.onLoginSuccess();
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
    onLoginSuccess() {
      this.getUsername();
    },
  },
};
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  background-color: #34495e;
  color: white;
  padding: 1rem 2rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
}

.navbar-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  max-width: 1200px;
}

.user-welcome {
  font-size: 1.5rem;
  margin: 0;
}

.logout-btn {
  background-color: rgb(79, 79, 79);
  color: white;
  border: solid 1px white;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

html,
body {
  height: 100%;
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
