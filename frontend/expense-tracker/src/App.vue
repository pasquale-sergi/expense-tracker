<template>
  <p v-if="isLogged">You are logged</p>
  <p v-else>You are not logged</p>
  <router-view></router-view>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "App",
  components: {},
  computed: {
    ...mapGetters(["isLogged"]),
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
