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
        <button @click="addIncome" class="btn-secondary">Add Income</button>
        <p>{{ message }}</p>
      </div>
    </div>
    <div class="income-history">
      <button @click="toggleIncomeHistory" class="btn-secondary">
        {{ showHistory ? "Hide" : "Show" }} Income History
      </button>

      <div class="popup-overlay" v-show="showHistory"></div>
      <!-- Popup Window -->
      <div class="popup" :class="{ show: showHistory }">
        <div class="popup-content">
          <span class="close" @click="toggleIncomeHistory">&times;</span>
          <income-history
            @close-popup="toggleIncomeHistory"
            :incomesHistory="groupedIncomes"
          ></income-history>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";
import IncomeHistory from "./IncomeHistory.vue";
export default {
  components: {
    IncomeHistory,
  },
  data() {
    return {
      incomeAmount: 0,
      incomeDescr: "",
      incomeDate: "",
      incomeRecord: [],
      incomeRecords: [],
      showHistory: false,
      message: "",
      incomesHistory: [],
    };
  },
  computed: {
    groupedIncomes() {
      const grouped = this.groupIncomesByMonth(this.incomesHistory);
      console.log("INFOOOO: ", this.incomesHistory);
      return Object.keys(grouped)
        .sort((a, b) => new Date(a) - new Date(b))
        .reduce((acc, key) => {
          acc[key] = grouped[key];
          return acc;
        }, {});
    },
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
    async toggleIncomeHistory() {
      this.showHistory = !this.showHistory;
      if (this.showHistory) {
        const response = await axios.get("http://localhost:8090/incomeHistory");
        this.incomesHistory = response.data;
        console.log("RESPONSE: ", response);
      }
    },
    formatDate(date) {
      return new Date(date).toLocaleDateString();
    },
    groupIncomesByMonth(incomes) {
      return incomes.reduce((acc, income) => {
        const date = new Date(income.Date);
        const year = date.getFullYear();
        const month = date.getMonth(); // Get month as a number (0-11)
        const key = new Date(year, month).toLocaleString("default", {
          month: "long",
          year: "numeric",
        });

        if (!acc[key]) {
          acc[key] = [];
        }

        acc[key].push(income);

        return acc;
      }, {});
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
.popup {
  display: none;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  border: 1px solid #ccc;
  padding: 40px;
  z-index: 1000;
  border-radius: 12px;
  max-height: 80%; /* Limit the height of the popup */
  overflow-y: auto; /* Enable vertical scrollbar when content exceeds the height */
  width: 80%;
  max-width: 800px;
}

.popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* Semi-transparent black overlay */
  z-index: 999; /* Ensure the overlay appears below the popup */
}

.popup-content {
  position: relative;
}

.close {
  position: absolute;
  top: 5px;
  right: 10px;
  cursor: pointer;
}

.show {
  display: block;
}
</style>