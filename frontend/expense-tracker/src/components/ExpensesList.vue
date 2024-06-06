<template>
  <div class="expenses">
    <div class="header">
      <h2>Expenses of The Month</h2>
    </div>
    <div class="expense-list">
      <div
        class="expense-item"
        v-for="expense in expenses"
        :key="expense.ExpenseID"
      >
        <div class="expense-details">
          <span class="amount">Amount: ${{ expense.Amount.toFixed(2) }}</span>
          <span class="description"
            >Description: {{ expense.Description }}</span
          >
          <span class="date">Date: {{ expense.Date }}</span>
        </div>
      </div>
    </div>
    <div class="exp-History">
      <button @click="getExpenseHistory" class="btn-secondary">
        {{ showHistory ? "Hide" : "Show" }} Expenses History
      </button>
      <div v-if="showHistory" class="history-list">
        <div
          class="expense-item"
          v-for="expense in expensesHistory"
          :key="expense.ExpenseID"
        >
          <div class="expense-item">
            <span class="amount">Amount: ${{ expense.Amount.toFixed(2) }}</span>
            <span class="description"
              >Description: {{ expense.Description }}</span
            >
            <span class="date">Date: {{ expense.Date }}</span>
          </div>
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
    this.intervalId = setInterval(this.getCurrentExpenses, 5000);
  },
  beforeUnmount() {
    // Clear the interval when the component is destroyed
    clearInterval(this.intervalId);
  },
  methods: {
    async getCurrentExpenses() {
      const response = await axios.get("http://localhost:8090/currentExpenses");
      this.expenses = response.data;
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
  max-width: 800px;
  margin: 20px auto;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.header {
  text-align: center;
  margin-bottom: 20px;
}

.expenses-list,
.history-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.expense-item {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
  background-color: #fff;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.expense-details {
  display: flex;
  flex-direction: column;
}

.amount,
.description,
.date {
  margin-bottom: 5px;
  font-size: 14px;
  color: #333;
}

.btn-secondary {
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  background-color: #6c757d;
  color: white;
  cursor: pointer;
  font-size: 16px;
  display: block;
  width: 100%;
  text-align: center;
  margin-top: 20px;
}

.btn-secondary:hover {
  opacity: 0.9;
}
</style>