<template>
  <div class="incomeSec">
    <h2>Income</h2>
    <div class="input">
      <label class="label" for="income">Income</label>
      <input
        class="input"
        type="number"
        placeholder="income amount"
        v-model="incomeAmount"
      />
    </div>
    <div class="input">
      <label class="label" for="income">Description</label>
      <input
        class="input"
        type="text"
        placeholder="description"
        v-model="incomeDescr"
      />
    </div>
    <div class="input">
      <label class="label" for="date">Date</label>
      <input
        class="input"
        type="text"
        placeholder="date"
        v-model="incomeDate"
      />
    </div>
    <div class="input">
      {{ message }}
      <button @click="addIncome">Add Income</button>
    </div>
  </div>
  <div class="incomeHistory">
    <div v-if="showIncomeHistory">
      <div v-for="income in incomeRecords" :key="income">
        Amount: {{ income.Amount }}, Description: {{ income.Description }},
        Date:
        {{ income.Date }}
      </div>
    </div>
    <button @click="incomeHistory">Show Income History</button>
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
.incomeSec {
  display: inline-block;
  border: solid 2px black;
  width: 50%;
  margin: auto;
}
.input {
  margin-top: 4px;
}

.label {
  margin-right: 10px;
}

.label {
  display: grid;
}
</style>