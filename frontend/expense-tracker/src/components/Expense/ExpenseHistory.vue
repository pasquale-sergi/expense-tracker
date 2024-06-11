<template>
  <div class="history-list">
    <div v-for="(expenses, month) in localExpensesHistory" :key="month">
      <h3>{{ month }}</h3>
      <table>
        <tr class="title-column">
          <th class="descr_content">Description</th>
          <th class="content">Amount</th>
          <th class="content">Date</th>
          <td class="delete">Delete</td>
        </tr>
        <tr
          class="columns-content"
          v-for="expense in expenses"
          :key="expense.ExpenseID"
        >
          <td class="descr_content">{{ expense.Description }}</td>
          <td class="content">${{ expense.Amount.toFixed(2) }}</td>
          <td class="content">
            {{ new Date(expense.Date).toLocaleDateString() }}
          </td>
          <td class="content">
            <button
              class="delete-exp"
              @click="deleteExpense(expense.Description, month)"
            >
              x
            </button>
          </td>
        </tr>
      </table>
    </div>
    <button class="close-btn" @click="closePopup">Close</button>
  </div>
</template>


<script>
import axios from "axios";
export default {
  props: {
    expensesHistory: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      localExpensesHistory: { ...this.expensesHistory },
    };
  },
  watch: {
    expensesHistory: {
      immediate: true,
      handler(newVal) {
        this.localExpensesHistory = { ...newVal };
      },
    },
  },
  methods: {
    closePopup() {
      this.$emit("close-popup");
    },
    formatDate(date) {
      return new Date(date).toLocaleDateString();
    },
    async deleteExpense(description, month) {
      try {
        const response = await axios.delete(
          `http://localhost:8090/deleteExpense/${description}`
        );

        // Check if the response indicates success (HTTP 2xx status codes)
        if (response.status >= 200 && response.status < 300) {
          // Remove the deleted expense from the local copy
          if (this.localExpensesHistory[month]) {
            this.localExpensesHistory[month] = this.localExpensesHistory[
              month
            ].filter((expense) => expense.Description !== description);

            // If the month has no more expenses, remove the month entry
            if (this.localExpensesHistory[month].length === 0) {
              delete this.localExpensesHistory[month];
            }
          }
        } else {
          console.error("Error response:", response);
          alert("Failed to delete expense");
        }
      } catch (error) {
        console.error("Error deleting expense:", error);
        alert("Failed to delete expense");
      }
    },
  },
};
</script>

<style scoped>
h3 {
  margin-top: 20px;
  margin-bottom: 20px;
  text-align: center;
}
.history-list {
  padding: 20px;
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

.columns-content {
  color: #333;
}

.descr_content {
  width: 40%;
}

.close-btn {
  margin-top: 30px;
  border-radius: 12px;
  padding-right: 10px;
  padding-left: 10px;
}

.delete-exp {
  padding: 2px 8px 2px 8px;
  border: none;
  border-radius: 100%;
  background-color: #ddd;
  color: black;
  cursor: pointer;
  margin-left: 12px;
}
</style>