<template>
  <h1>Summary Section</h1>
  <data-table :items="items"></data-table>
</template>
  
  <script setup>
import { onMounted, ref } from "vue";
import axios from "axios";
import DataTable from "./DataTable.vue";

// Declare reactive items array
const items = ref([]);

// Fetch data on component mount
onMounted(async () => {
  try {
    const response = await axios.get("http://localhost:8090/currentExpenses");
    items.value = response.data;
    console.log("REsponse: ", response);
    console.log("ITems: ", items);
  } catch (error) {
    console.error("Error fetching data:", error);
  }
});
</script>
  
  <script>
export default {
  components: {
    DataTable,
  },
};
</script>