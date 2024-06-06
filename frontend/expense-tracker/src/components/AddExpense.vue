<template>
  <div class="addExp">
    <hr class="vertical" />
    <div id="new">
      <h2>Add new Expense</h2>
      <div class="inputs">
        <div class="inputitem">
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
        <div class="inputitem">
          <label class="label">Category:</label>
          <select name="category" v-model="categoryChoice">
            <option v-for="category in categoryOptions" :key="category">
              {{ category.Name }}
            </option>
          </select>
        </div>

        <div class="inputitem">
          <label class="label">Date:</label>
          <input type="text" placeholder="date" v-model="expenseDate" />
        </div>
      </div>
      <button @click="addExpense">Add Expense</button>
      <p>{{ message }}</p>
      <button @click="showPop">Add Category</button>
      <add-category v-if="showPopUp"></add-category>
    </div>
  </div>
</template>


<script>
import axios from "axios";
import AddCategory from "./AddCategory.vue";
export default {
  components: {
    AddCategory,
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
  },
};
</script>


<style scoped>
.addExp {
  border: solid 2px black;
  width: 50%;
  margin: auto;
}
.inputitem {
  align-content: center;
  align-items: center;
  text-align: center;
}

.label {
  display: grid;
}

button {
  margin-top: 5px;
}
</style>
