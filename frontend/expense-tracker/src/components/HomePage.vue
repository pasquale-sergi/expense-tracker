<template>
  <h2>Hello Page</h2>
  <div class="balance">
    <user-balance></user-balance>
  </div>

  <div class="expenses">
    <expenses-list v-if="isLogged"></expenses-list>
  </div>
  <div class="addExpense">
    <add-expense></add-expense>
  </div>
  <div class="Income">
    <income-section></income-section>
  </div>
</template>

<script>
import ExpensesList from "./ExpensesList.vue";
import AddExpense from "./AddExpense.vue";
import IncomeSection from "./IncomeSection.vue";
import UserBalance from "./UserBalance.vue";
import { mapGetters } from "vuex";

export default {
  components: {
    ExpensesList,
    AddExpense,
    IncomeSection,
    UserBalance,
  },
  computed: {
    ...mapGetters(["isLogged"]),
  },
};
</script>