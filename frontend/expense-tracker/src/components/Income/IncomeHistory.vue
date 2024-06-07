<template>
  <div class="history-list">
    <div v-for="(incomes, month) in incomesHistory" :key="month">
      <h3>{{ month }}</h3>
      <table>
        <tr class="title-column">
          <th class="descr_content">Description</th>
          <th class="content">Amount</th>
          <th class="content">Date</th>
        </tr>
        <tr class="columns-content" v-for="income in incomes" :key="income">
          <td class="descr_content">{{ income.Description }}</td>
          <td class="content">${{ income.Amount.toFixed(2) }}</td>
          <td class="content">
            {{ new Date(income.Date).toLocaleDateString() }}
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
    incomesHistory: {
      type: Object,
      required: true,
    },
  },
  created() {
    console.log("Income History Prop:", this.incomesHistory); // Log the received prop
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