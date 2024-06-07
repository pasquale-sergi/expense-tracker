<template>
  <div>
    <table class="summary-table">
      <tr class="title-column">
        <th class="content">Amount</th>
        <th class="content">Category</th>
      </tr>
      <tr class="columns-content" v-for="item in items" :key="item">
        <td class="content">${{ item.total_amount }}</td>
        <td class="content">
          {{ item.category_name === "" ? "No category" : item.category_name }}
        </td>
      </tr>
    </table>
  </div>
</template>

<script setup>
import { defineProps } from "vue";

// Define props using defineProps macro
defineProps({
  items: {
    type: Array,
    required: true,
  },
});

// Define reactive variable to store the current month
// const months = [
//   "January",
//   "February",
//   "March",
//   "April",
//   "May",
//   "June",
//   "July",
//   "August",
//   "September",
//   "October",
//   "November",
//   "December",
// ];

// const d = new Date();

// let month_name = months[d.getMonth()];
</script>

<style scoped>
.title-column {
  margin-left: 10px;
  border: solid 2px black;
}
.descr_content {
  padding-right: 50px;
}
.content {
  padding: 10px;
}

.summary-table {
  margin-left: 230px;
  margin-top: 20px;
  border: solid 2px black;
}

.changemonth-btn {
  margin-top: 30px;
  padding: 4px;
  border-radius: 12px;
}
</style>
