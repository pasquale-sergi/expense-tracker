<template>
  <div class="screen-1">
    <h2 class="title">Register</h2>
    <form @submit.prevent="register">
      <div class="username">
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
      <div class="email">
        <label for="email">Email</label>
        <input
          class="input"
          type="email"
          placeholder="email"
          id="email"
          v-model="email"
          required
        />
      </div>
      <div class="password">
        <label for="password">Password</label>
        <input
          class="input"
          type="password"
          placeholder="password"
          id="password"
          v-model="password"
          required
        />
      </div>
      <div class="password">
        <label for="confirmPassword">Confirm Password</label>
        <input
          class="input"
          type="password"
          placeholder="confirm password"
          id="confirmPassword"
          v-model="confirmPassword"
          required
        />
      </div>
      <p class="message">{{ message }}</p>
      <button class="login" type="submit">Register</button>
      <div class="footer">
        <a href="#" @click.prevent="switchToLogin"
          >Already have an account? Login</a
        >
      </div>
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
      confirmPassword: "",
      message: "",
    };
  },
  methods: {
    ...mapActions(["registerUser"]),

    switchToLogin() {
      this.$emit("switch-view", "login");
    },

    register() {
      if (this.password === this.confirmPassword) {
        this.registerUser({
          username: this.username,
          email: this.email,
          password: this.password,
        });

        this.message = "You are now registered! Please Log In.";
        this.username = "";
        this.email = "";
        this.password = "";
        this.confirmPassword = "";
      } else {
        this.message = "Passwords are not matching. Please try again";
      }
    },
  },
};
</script>

<style scoped>
/* Reuse your styles from the login component */
.title {
  margin-top: 20px;
  margin-bottom: 10px;
}
.screen-1 {
  width: 27%;
  display: block;
  border: solid 3px black;
  border-radius: 20px;
  justify-content: center;
  text-align: center;
}
.email,
.username {
  background: hsl(0deg 0% 100%);
  box-shadow: 0 0 2em hsl(231deg 62% 94%);
  padding: 1em;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
  border-radius: 20px;
  color: hsl(0deg 0% 30%);
  margin-top: 10px;
  width: 60%;
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

  width: 60%;
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
  margin: 10px 30px 10px 173px;
}

button {
  cursor: pointer;
}

.input {
  border-radius: 12px;
  text-align: center;
}
.username {
  margin-bottom: 30px;
}

.message {
  margin-top: 15px;
}
</style>





