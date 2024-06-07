<template>
  <div>
    <h2>Register</h2>
    <form @submit.prevent="register">
      <div>
        <label for="name">Username</label>
        <input
          type="text"
          id="name"
          v-model="username"
          placeholder="enter your username"
          required
        />
      </div>
      <div>
        <label for="email">Email</label>
        <input
          type="text"
          id="email"
          v-model="email"
          placeholder="enter your email"
          required
        />
      </div>
      <div>
        <label for="password">Password</label>
        <input
          type="password"
          id="password"
          v-model="password"
          placeholder="enter your password"
          required
        />
      </div>
      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  data() {
    return {
      username: "",
      email: "",
      password: "",
    };
  },
  methods: {
    ...mapActions(["registerUser"]),
    register() {
      this.registerUser({
        username: this.username,
        email: this.email,
        password: this.password,
      });
    },
  },
};
</script>