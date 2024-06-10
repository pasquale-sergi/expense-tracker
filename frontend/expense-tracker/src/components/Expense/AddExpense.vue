<template>
  <div class="add-exp">
    <div class="form-container">
      <h2>Add new Expense</h2>
      <div class="inputs">
        <div class="input-item">
          <label class="label">Description:</label>
          <input
            id="name"
            type="text"
            placeholder="description"
            v-model="expenseDescr"
          />
        </div>
        <div class="inputitem">
          <label class="label">Amount:</label>
          <input
            id="amount"
            type="number"
            placeholder="amount"
            v-model="expenseAmount"
          />
        </div>
        <div class="input-item">
          <label class="label">Category:</label>
          <select name="category" v-model="categoryChoice">
            <option v-for="category in categoryOptions" :key="category">
              {{ category.Name }}
            </option>
          </select>
        </div>

        <div class="input-item">
          <label class="label">Date:</label>
          <input type="text" placeholder="date" v-model="expenseDate" />
        </div>
      </div>
      <button @click="addExpense">Add Expense</button>
      <p>{{ message }}</p>
      <button @click="showPop">Add Category</button>
      <add-category v-if="showPopUp"></add-category>
      <div class="fixed-exp">
        <button class="fixed-exp" @click="toggleFixedForm">
          Set Fixed Expenses
        </button>

        <fixed-expenses
          v-if="showFixedForm"
          @close-popup="toggleFixedForm"
          :categoryOptions="categoryOptions"
        ></fixed-expenses>
      </div>
    </div>
  </div>
</template>


<script>
import axios from "axios";
import AddCategory from "./AddCategory.vue";
import FixedExpenses from "./FixedExpenses.vue";
export default {
  components: {
    AddCategory,
    FixedExpenses,
  },
  data() {
    return {
      expenseAmount: 0,
      expenseDescr: "",
      expenseDate: "",
      expense: [],
      message: "",
      categoryOptions: [],
      categoryChoice: "",
      showPopUp: false,
      showFixedForm: false,
    };
  },
  created() {
    this.getCategories();
  },
  methods: {
    async addExpense() {
      this.expense = {
        Amount: this.expenseAmount,
        Categoryname: this.categoryChoice,
        Description: this.expenseDescr,
        Date: this.expenseDate,
      };
      try {
        await axios.post("http://localhost:8090/addExpense", this.expense);
        this.message = "Expense added to the list";
      } catch (err) {
        console.log("Error adding the expense: ", err);
        this.message = "Expense not added to the list";
      }
    },
    async getCategories() {
      try {
        const response = await axios.get("http://localhost:8090/categories");
        this.categoryOptions = response.data;
        console.log(response);
      } catch (err) {
        console.log("Error getting the categories list: ", err);
      }
    },
    showPop() {
      this.showPopUp = !this.showPopUp;
      console.log(this.showPopUp);
    },
    toggleFixedForm() {
      this.showFixedForm = !this.showFixedForm;
    },
  },
};
</script>


<style scoped>
.add-exp {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  max-width: 600px;
  margin: auto;
}

.form-container {
  width: 100%;
}

h2 {
  text-align: center;
  color: #333;
}

.inputs {
  display: flex;
  flex-direction: column;
  gap: 15px;
  margin-bottom: 20px;
}

.input-item {
  display: flex;
  flex-direction: column;
}

.label {
  margin-bottom: 5px;
  font-weight: bold;
  color: #555;
}

input,
select {
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

.fixed-exp {
  margin-top: 10px;
}
</style>