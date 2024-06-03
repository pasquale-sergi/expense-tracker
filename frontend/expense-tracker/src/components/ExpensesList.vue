<template>
  <div>hi</div>

  <div v-for="expense in expenses" :key="expense.ExpenseID">
    Amount: {{ expense.Amount }}, Description: {{ expense.Description }}, Date:
    {{ expense.Date }}
  </div>
</template>


<script>
import axios from "axios";
export default {
  data() {
    return {
      expenses: [],
    };
  },
  created() {
    this.expensesList();
  },
  methods: {
    async expensesList() {
      const userID = 2;

      const response = await axios.get(
        "http://localhost:8090/expenses",
        userID
      );
      this.expenses = response.data;
      console.log(response);
    },
  },
};
</script>