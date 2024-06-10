<template>
  <div class="auth-form">
    <login-user v-if="isLoginView" @switch-view="switchView" />
    <register-user v-else @switch-view="switchView" />
  </div>
</template>
  
  <script>
import LoginUser from "./LoginUser.vue";
import RegisterUser from "./RegisterUser.vue";

export default {
  data() {
    return {
      isLoginView: true,
    };
  },
  components: {
    LoginUser,
    RegisterUser,
  },
  methods: {
    switchView(view) {
      this.isLoginView = view === "login";
    },
  },
};
</script>

<style scoped>
.auth-form {
  width: auto;
  margin: auto;
  margin-top: 15%;
  display: flex;
  justify-content: center;
  align-content: center;
  align-items: center;
}
</style>