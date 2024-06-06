<template>
  <h2>Current Balance</h2>
  <div class="balance">$ {{ parseFloat(balance) }}</div>
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