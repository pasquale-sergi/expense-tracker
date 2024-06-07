<template>
  <div>
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div>
        <label for="username">Username</label>
        <input
          type="text"
          placeholder="username"
          id="username"
          v-model="username"
          required
        />
      </div>
      <div>
        <label for="passord">Password</label>
        <input
          type="password"
          placeholder="password"
          id="password"
          v-model="password"
          required
        />
      </div>
      <button type="submit">Login</button>
      <router-link to="/register"
        >Are you not registered? Do it Right Now</router-link
      >
    </form>
    <p>{{ loginMessage }}</p>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  computed: {
    ...mapGetters(["loginMessage"]),
  },
  methods: {
    ...mapActions(["loginUser"]),
    login() {
      this.loginUser({
        username: this.username,
        password: this.password,
      });
    },
  },
};
</script>