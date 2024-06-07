<template>
  <div class="history-list">
    <div v-for="(expenses, month) in expensesHistory" :key="month">
      <h3>{{ month }}</h3>
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
          <td class="content">
            {{ new Date(expense.Date).toLocaleDateString() }}
          </td>
        </tr>
      </table>
    </div>
    <button class="close-btn" @click="closePopup">Close</button>
  </div>
</template>


<script>
export default {
  props: {
    expensesHistory: {
      type: Object,
      required: true,
    },
  },
  created() {
    console.log("Expenses History Prop:", this.expensesHistory); // Log the received prop
  },
  methods: {
    closePopup() {
      this.$emit("close-popup");
    },
    formatDate(date) {
      return new Date(date).toLocaleDateString();
    },
  },
};
</script>

<style scoped>
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
</style>