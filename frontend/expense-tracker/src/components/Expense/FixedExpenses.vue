<template>
  <div class="modal-overlay" @click.self="closePopup">
    <div class="modal-content">
      <span class="close-button" @click="closePopup">&times;</span>
      <h2>Fixed Expenses</h2>
      <div class="form">
        <label for="name">Expense Name</label>
        <input type="text" placeholder="name" required v-model="name" />
        <label for="amount">Amount</label>
        <input type="number" placeholder="amount" required v-model="amount" />
        <label for="category">Category</label>
        <select name="category" id="category" v-model="categoryName">
          <option v-for="category in categoryOptions" :key="category">
            {{ category.Name }}
          </option>
        </select>
      </div>
      <p>{{ message }}</p>
      <button @click="setFixedExpenses">Add Expense</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  props: ["categoryOptions"],
  data() {
    return {
      name: "",
      amount: "",
      categoryName: "",
      expense: [],
      message: "",
    };
  },
  methods: {
    async setFixedExpenses() {
      try {
        this.expense = {
          Name: this.name,
          Amount: this.amount,
          Categoryname: this.categoryName,
        };
        await axios.post("http://localhost:8090/fixedExpenses", this.expense);

        this.message = "Fixed Expense added.";
      } catch (error) {
        console.log(error);
        this.message = "Something went wrong.";
      }
    },
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  width: 500px;
  max-width: 90%;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  position: relative;
  text-align: center;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  cursor: pointer;
  font-size: 20px;
  color: #aaa;
}

.close-button:hover {
  color: #000;
}

.form {
  display: flex;
  flex-direction: column;
}

label {
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
  margin-bottom: 15px;
  text-align: center;
}

button {
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  background-color: #6e6e6e;
  color: white;
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
</style>