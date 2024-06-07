  <template>
  <nav class="navbar" v-if="isLogged">
    <p class="user-welcome" v-if="isLogged">
      Welcome Back {{ capitalizeFirstLetter(username) }}
    </p>

    <button v-if="isLogged" @click="logoutUser" class="logout-btn">
      Logout
    </button>
  </nav>

  <login-user v-if="!isLogged" class="login-box"></login-user>

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

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh; /* Ensures the container takes the full height of the viewport */
  margin: 0; /* Remove any default margin */
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  width: 100%;
  padding: 20px;
  border: solid 1px black;
  margin-bottom: 20px;
  border-radius: 8px;
}

.user-welcome {
  font-size: 25px;
}

.login-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: solid 1px black;
  padding: 20px;
  border-radius: 25px;
  width: 450px; /* Adjust the width as needed */
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); /* Optional: add a shadow for better visual appearance */
  background-color: #fff; /* Optional: add a background color */
  margin: auto; /* Center the login box */
  margin-top: 10%;
}

html,
body {
  height: 100%; /* Ensure the body takes up the full height of the viewport */
  margin: 0; /* Remove default margin */
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
