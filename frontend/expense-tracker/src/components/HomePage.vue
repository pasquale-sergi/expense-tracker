<template>
  <h2>Hello Page</h2>

  <div class="expenses">
    <expenses-list v-if="isLogged"></expenses-list>
  </div>
  <div class="addExpense">
    <add-expense></add-expense>
  </div>
</template>

<script>
import ExpensesList from "./ExpensesList.vue";
import AddExpense from "./AddExpense.vue";
import { mapGetters } from "vuex";

export default {
  components: {
    ExpensesList,
    AddExpense,
  },
  computed: {
    ...mapGetters(["isLogged"]),
  },
};
</script>