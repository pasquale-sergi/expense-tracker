<template>
  <div class="screen-1">
    <h2 class="title">Login</h2>
    <form @submit.prevent="login">
      <div class="email">
        <label for="username">Username</label>
        <input
          class="input"
          type="text"
          placeholder="username"
          id="username"
          v-model="username"
          required
        />
      </div>
      <div class="password">
        <label for="passord">Password</label>
        <input
          class="input"
          type="password"
          placeholder="password"
          id="password"
          v-model="password"
          required
        />
      </div>
      <button class="login" type="submit">Login</button>
      <div class="footer">
        <router-link to="/register"
          >Are you not registered? Do it Right Now</router-link
        >
      </div>
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

<style scoped>
.title {
  margin-top: 20px;
  margin-bottom: 10px;
}
.screen-1 {
  width: 20%;
  display: block;
  border: solid 3px black;
  border-radius: 20px;
  justify-content: center;
  text-align: center;
}
.email {
  background: hsl(0deg 0% 100%);
  box-shadow: 0 0 2em hsl(231deg 62% 94%);
  padding: 1em;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
  border-radius: 20px;
  color: hsl(0deg 0% 30%);
  margin-top: 10px;
  width: 80%;
  margin: auto;
}
.password {
  background: hsl(0deg 0% 100%);
  box-shadow: 0 0 2em hsl(231deg 62% 94%);
  padding: 1em;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
  border-radius: 20px;
  color: hsl(0deg 0% 30%);

  width: 80%;
  margin: auto;
  margin-top: 30px;
}
.login {
  padding: 10px;
  background: hsl(233deg 36% 38%);
  color: hsl(0 0 100);
  border: none;
  border-radius: 30px;
  font-weight: 600;
  margin-top: 10px;
  padding: 10px 25px 10px 25px;
}
.footer {
  display: flex;
  font-size: 0.7em;
  color: hsl(0deg 0% 37%);
  gap: 14em;
  padding-bottom: 10em;
  margin-top: 10px;
  width: 80%;
  margin: 10px 30px 10px 130px;
}

button {
  cursor: pointer;
}

.input {
  border-radius: 12px;
  text-align: center;
}
</style>