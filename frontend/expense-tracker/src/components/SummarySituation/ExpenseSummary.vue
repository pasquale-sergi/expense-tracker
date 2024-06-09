<template>
  <h1>Summary Section</h1>
  <div class="box">
    <div class="chart">
      <PieChart :items="items" />
    </div>

    <div class="category-amount">
      <h3>{{ selectedMonth || getCurrentMonth }}</h3>
      <div class="table">
        <data-table :items="items"></data-table>
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
  </div>
</template>

<script>
import axios from "axios";
import DataTable from "./DataTable.vue";

import PieChart from "./PieChart.vue";

export default {
  components: {
    DataTable,
    PieChart,
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
      return this.months[d.getMonth()];
    },
  },
  mounted() {
    this.fetchSummaryData();
    this.selectedMonth = this.getCurrentMonth();
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
  display: grid;
  grid-template-columns: 1fr 1fr; /* Two columns */
  gap: 20px; /* Spacing between columns */
  text-align: center;
  background-color: #f5f5f5;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  max-width: 80%;
  margin: auto;
  height: auto;
  padding-bottom: 30px;
  padding-top: 30px;
}

.chart {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  grid-column: 1; /* Chart on the left */
  grid-row: 1; /* Span 2 rows */
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.category-amount {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  grid-column: 2;
  grid-row: 1;
}

.table {
  padding-top: 25px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

h1 {
  text-align: center;
  padding: 20px;
}
</style>