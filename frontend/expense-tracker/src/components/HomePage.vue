<template>
  <div class="balance">
    <user-balance></user-balance>
  </div>
  <div class="container">
    <div class="expenses">
      <expenses-list v-if="isLogged"></expenses-list>
    </div>
    <div class="add-expense">
      <add-expense></add-expense>
    </div>
    <div class="income">
      <income-section></income-section>
    </div>
    <div class="summary">
      <expense-summary></expense-summary>
    </div>
  </div>
</template>

<script>
import ExpensesList from "./Expense/ExpensesList.vue";
import AddExpense from "./Expense/AddExpense.vue";
import IncomeSection from "./Income/IncomeSection.vue";
import UserBalance from "./User/UserBalance.vue";
import { mapGetters } from "vuex";
import ExpenseSummary from "./Expense/ExpenseSummary.vue";

export default {
  components: {
    ExpensesList,
    AddExpense,
    IncomeSection,
    UserBalance,
    ExpenseSummary,
  },
  computed: {
    ...mapGetters(["isLogged"]),
  },
};
</script>

<style scoped>
.container {
  display: grid;

  grid-template-rows: auto auto;
  gap: 20px;
  padding: 20px;
  max-width: 2000px;
  margin: 0 auto;
  grid-template-columns: 1fr 1fr 1fr;
}

.balance {
  grid-column: 2;
  grid-row: 2;
  text-align: center;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 50%;
  margin: auto;
}

.expenses,
.add-expense,
.income {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.expenses {
  grid-column: 1;
  grid-row: 1;
  height: auto;
}

.add-expense {
  grid-column: 2;
  grid-row: 1;
}

.income {
  grid-column: 3;
  grid-row: 1;
}
</style>