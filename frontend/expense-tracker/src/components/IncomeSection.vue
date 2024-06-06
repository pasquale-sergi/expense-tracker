<template>
  <div class="income-section">
    <div class="form-container">
      <h2>Income</h2>
      <div class="input-group">
        <label class="label" for="income">Income</label>
        <input
          class="input"
          type="number"
          placeholder="income amount"
          v-model="incomeAmount"
        />
      </div>
      <div class="input-group">
        <label class="label" for="income">Description</label>
        <input
          class="input"
          type="text"
          placeholder="description"
          v-model="incomeDescr"
        />
      </div>
      <div class="input-group">
        <label class="label" for="date">Date</label>
        <input
          class="input"
          type="text"
          placeholder="date"
          v-model="incomeDate"
        />
      </div>
      <div class="input-group">
        <button @click="addIncome" class="btn-primary">Add Income</button>
        <p>{{ message }}</p>
      </div>
    </div>
    <div class="income-history">
      <button @click="incomeHistory" class="btn-secondary">
        {{ showIncomeHistory ? "Hide" : "Show" }} Income History
      </button>
      <div v-if="showIncomeHistory" class="history-list">
        <div v-for="income in incomeRecords" :key="income">
          Amount: ${{ income.Amount.toFixed(2) }}, Description:
          {{ income.Description }}, Date:
          {{ income.Date }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      incomeAmount: 0,
      incomeDescr: "",
      incomeDate: "",
      incomeRecord: [],
      incomeRecords: [],
      showIncomeHistory: false,
      message: "",
    };
  },
  methods: {
    async addIncome() {
      this.incomeRecord = {
        Amount: this.incomeAmount,
        Description: this.incomeDescr,
        Date: this.incomeDate,
      };
      try {
        await axios.post("http://localhost:8090/addIncome", this.incomeRecord);
        console.log("income record added correctly");
        this.message = "Income Record added to the list";
      } catch (err) {
        console.log("Error adding the income record: ", err);
        this.message = "Income Record not added to the list";
      }
    },
    async incomeHistory() {
      try {
        this.showIncomeHistory = !this.showIncomeHistory;
        const response = await axios.get("http://localhost:8090/incomeHistory");
        this.incomeRecords = response.data;
        console.log("incomes history everything is good");
      } catch (err) {
        console.log("error while showing income history: ", err);
      }
    },
  },
};
</script>

<style scoped>
.income-section {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.form-container {
  width: 100%;
}

h2 {
  text-align: center;
  color: #333;
  margin-bottom: 20px;
}

.input-group {
  display: flex;
  flex-direction: column;
  margin-bottom: 15px;
}

.label {
  margin-bottom: 5px;
  font-weight: bold;
  color: #555;
}

.input {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 16px;
}

input[type="date"] {
  appearance: none;
  -webkit-appearance: none;
}

button {
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}

.btn-primary {
  background-color: #007bff;
  color: white;
  margin-top: 10px;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
  margin-top: 10px;
  display: block;
  width: 100%;
  text-align: center;
}

button:hover {
  opacity: 0.9;
}

p {
  text-align: center;
  color: #28a745;
  font-weight: bold;
  margin-top: 10px;
}

.income-history {
  margin-top: 20px;
}

.history-list {
  margin-top: 20px;
  border-top: 1px solid #ddd;
  padding-top: 10px;
}

.history-list div {
  margin-bottom: 10px;
  font-size: 14px;
}
</style>