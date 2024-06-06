<template>
  <h2>Current Balance</h2>
  <div class="balance">$ {{ parseFloat(balance).toFixed(2) }}</div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      balance: null,
    };
  },
  created() {
    this.getBalance();
    this.intervalId = setInterval(this.getBalance, 5000);
  },
  beforeUnmount() {
    // Clear the interval when the component is destroyed
    clearInterval(this.intervalId);
  },

  methods: {
    async getBalance() {
      try {
        const response = await axios.get("http://localhost:8090/balance");
        this.balance = parseFloat(response.data);
      } catch (err) {
        console.log("Error getting balance", err);
      }
    },
  },
};
</script>

<style scoped>
.balance {
  font-size: 50px;
}
</style>