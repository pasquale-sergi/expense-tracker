<template>
  <div>
    <h2>Pie Chart</h2>
    <div class="chart-container">
      <Pie :data="chartData" :options="chartOptions" />
    </div>
  </div>
</template>

<script>
import { Pie } from "vue-chartjs";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  ArcElement,
  CategoryScale,
  LinearScale,
} from "chart.js";

ChartJS.register(
  Title,
  Tooltip,
  Legend,
  ArcElement,
  CategoryScale,
  LinearScale
);

export default {
  props: ["items"],
  name: "PieChart",
  components: {
    Pie,
  },
  computed: {
    chartData() {
      const defaultColors = [
        "#FF6384",
        "#4BC0C0",
        "#FFCE56",
        "#36A2EB",
        "#9966FF",
        "#FF9F40",
      ];
      if (!this.items) {
        return { labels: ["No Data"], datasets: [100] }; // Return empty data if items is null or undefined
      }
      return {
        labels: this.items.map((item) =>
          item.category_name === "" ? "No category" : item.category_name
        ),
        datasets: [
          {
            label: "Monthly Summary",
            backgroundColor: defaultColors.slice(0, this.items.length),
            data: this.items.map((item) => item.total_amount),
          },
        ],
      };
    },
    chartOptions() {
      return {
        responsive: true,
        plugins: {
          legend: {
            position: "top",
          },
          title: {
            display: true,
            text: "Monthly Summary Pie Chart",
          },
        },
      };
    },
  },
  watch: {
    items: {
      handler() {
        this.updateChart();
      },
      deep: true,
    },
  },
  methods: {
    updateChart() {
      if (this.$refs.pieChart && this.$refs.pieChart.chart) {
        this.$refs.pieChart.chart.update();
      }
    },
  },
  mounted() {
    // Ensure the chart is updated after the component is mounted
    this.updateChart();
  },
};
</script>

<style scoped>
.chart-container {
  position: relative;
  max-width: 500px;
  margin: 0 auto;
  height: 400px; /* Ensure the container has a height */
  padding: 20px;
}

canvas {
  max-width: 100%;
  height: auto !important;
}
</style>
