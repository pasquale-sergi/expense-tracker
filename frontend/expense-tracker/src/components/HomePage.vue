<template>
  <h2>Hello Page</h2>

  <div class="expenses">
    <expenses-list v-if="isLogged"></expenses-list>
  </div>
  <router-link to="/register">Register page</router-link>
</template>

<script>
import ExpensesList from "./ExpensesList.vue";
import { mapGetters } from "vuex";

export default {
  components: {
    ExpensesList,
  },
  computed: {
    ...mapGetters(["isLogged"]),
  },
};
</script>