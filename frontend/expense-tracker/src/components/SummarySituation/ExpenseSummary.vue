<template>
  <div class="box">
    <h1>Summary Section</h1>
    <div class="category-amount">
      <h3>{{ selectedMonth === null ? getCurrentMonth() : selectedMonth }}</h3>
      <data-table :items="items" class="table"></data-table>
    </div>
    <select
      v-model="selectedMonth"
      @change="fetchSummaryData"
      class="monthChoice"
    >
      <option v-for="month in months" :key="month" :value="month">
        {{ month }}
      </option>
    </select>
  </div>
</template>

<script>
import axios from "axios";
import DataTable from "./DataTable.vue";

export default {
  components: {
    DataTable,
  },
  data() {
    return {
      items: [],
      months: [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
      ],
      selectedMonth: null,
    };
  },
  methods: {
    async fetchSummaryData() {
      try {
        const response = await axios.get(
          "http://localhost:8090/summaryByCategoryMonth",
          {
            params: { month: this.selectedMonth },
          }
        );
        this.items = response.data;
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    getCurrentMonth() {
      const d = new Date();
      this.selectedMonth = this.months[d.getMonth()];
    },
  },
  mounted() {
    this.fetchSummaryData();
  },
};
</script>

<style scoped>
.title-column {
  margin-left: 10px;
  border: solid 2px black;
}

.monthChoice {
  margin-top: 30px;
  border-radius: 15px;
  padding: 5px;
  text-align: center;
}

.box {
  display: block;
  text-align: center;
  background-color: #f5f5f5;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  max-width: 600px;
  margin: auto;
  height: auto;
  padding-bottom: 30px;
  padding-top: 30px;
}

.table {
}
</style>