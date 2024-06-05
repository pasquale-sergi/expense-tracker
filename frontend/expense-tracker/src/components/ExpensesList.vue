<template>
  <div class="expenses">
    <div class="expOfMonth"></div>
    <div>Expenses of The Month</div>

    <div v-for="expense in expenses" :key="expense.ExpenseID">
      Amount: {{ expense.Amount }}, Description: {{ expense.Description }},
      Date:
      {{ expense.Date }}
    </div>
    <div class="expHistory">
      <button @click="getExpenseHistory">Show Expenses History</button>
      <div v-if="showHistory">
        <div v-for="expense in expensesHistory" :key="expense.ExpenseID">
          Amount: {{ expense.Amount }}, Description: {{ expense.Description }},
          Date:
          {{ expense.Date }}
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import axios from "axios";
export default {
  data() {
    return {
      expenses: [],
      expensesHistory: [],
      showHistory: false,
    };
  },
  created() {
    this.getCurrentExpenses();
  },
  methods: {
    async getCurrentExpenses() {
      const response = await axios.get("http://localhost:8090/currentExpenses");
      this.expenses = response.data;
      console.log("RESPONSE: ", response);
    },
    async getExpenseHistory() {
      this.showHistory = !this.showHistory;
      const response = await axios.get("http://localhost:8090/expensesHistory");
      this.expensesHistory = response.data;
    },
  },
};
</script>

<style scoped>
.expenses {
  border: solid 2px black;
  margin: auto;
  width: 50%;
  display: block;
  align-content: center;
  align-items: center;
}
</style>