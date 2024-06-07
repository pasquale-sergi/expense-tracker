<template>
  <div class="expenses">
    <div class="header">
      <h2>Expenses of The Month</h2>
    </div>
    <div class="expense-list">
      <table>
        <tr class="title-column">
          <th class="descr_content">Description</th>
          <th class="content">Amount</th>
          <th class="content">Date</th>
        </tr>
        <tr
          class="columns-content"
          v-for="expense in expenses"
          :key="expense.ExpenseID"
        >
          <td class="descr_content">{{ expense.Description }}</td>
          <td class="content">${{ expense.Amount.toFixed(2) }}</td>
          <td class="content">{{ formatDate(expense.Date) }}</td>
        </tr>
      </table>
    </div>
    <div class="exp-History">
      <button @click="toggleExpenseHistory" class="btn-secondary">
        {{ showHistory ? "Hide" : "Show" }} Expenses History
      </button>
      <div class="popup-overlay" v-show="showHistory"></div>
      <!-- Popup Window -->
      <div class="popup" :class="{ show: showHistory }">
        <div class="popup-content">
          <span class="close" @click="toggleExpenseHistory">&times;</span>
          <expense-history
            @close-popup="toggleExpenseHistory"
            :expensesHistory="groupedExpenses"
          ></expense-history>
        </div>
      </div>
    </div>
  </div>
</template>



<script>
import axios from "axios";
import ExpenseHistory from "./ExpenseHistory.vue";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";
export default {
  components: {
    ExpenseHistory,
  },
  data() {
    return {
      expenses: [],
      expensesHistory: [],
      showHistory: false,
      expenseDate: "",
    };
  },
  computed: {
    groupedExpenses() {
      const grouped = this.groupExpensesByMonth(this.expensesHistory);
      return Object.keys(grouped)
        .sort((a, b) => new Date(a) - new Date(b))
        .reduce((acc, key) => {
          acc[key] = grouped[key];
          return acc;
        }, {});
    },
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
      try {
        const response = await axios.get(
          "http://localhost:8090/currentExpenses"
        );
        this.expenses = response.data;
      } catch (error) {
        console.error("Error fetching current expenses:", error);
      }
    },
    async toggleExpenseHistory() {
      this.showHistory = !this.showHistory;
      if (this.showHistory) {
        console.log(this.showHistory);
        const response = await axios.get(
          "http://localhost:8090/expensesHistory"
        );
        this.expensesHistory = response.data;

        console.log(response);
      }
    },
    formatDate(date) {
      return new Date(date).toLocaleDateString();
    },
    groupExpensesByMonth(expenses) {
      return expenses.reduce((acc, expense) => {
        const date = new Date(expense.Date);
        const year = date.getFullYear();
        const month = date.getMonth(); // Get month as a number (0-11)
        const key = new Date(year, month).toLocaleString("default", {
          month: "long",
          year: "numeric",
        });

        if (!acc[key]) {
          acc[key] = [];
        }

        acc[key].push(expense);

        return acc;
      }, {});
    },
  },
};
</script>

<style scoped>
.expenses {
  padding: 20px;
}

.header {
  margin-bottom: 20px;
}

.expense-list {
  max-height: 400px; /* Adjust this value as needed */
  overflow-y: auto;
  margin-bottom: 20px;
  border: 1px solid #ddd;
  border-radius: 5px;
  background-color: #fff;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  padding: 10px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.title-column {
  background-color: #f4f4f4;
}

.content {
  color: #333;
}

.descr_content {
  width: 40%;
}

.exp-History {
  margin-top: 20px;
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

.popup {
  display: none;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  border: 1px solid #ccc;
  padding: 40px;
  z-index: 1000;
  border-radius: 12px;
  max-height: 80%; /* Limit the height of the popup */
  overflow-y: auto; /* Enable vertical scrollbar when content exceeds the height */
  width: 80%;
  max-width: 800px;
}

.popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* Semi-transparent black overlay */
  z-index: 999; /* Ensure the overlay appears below the popup */
}

.popup-content {
  position: relative;
}

.close {
  position: absolute;
  top: 5px;
  right: 10px;
  cursor: pointer;
}

.show {
  display: block;
}
</style>