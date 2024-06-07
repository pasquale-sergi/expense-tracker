<template>
  <h2>Add a new category</h2>
  <div class="input">
    <label for="name">Category Name</label>
    <input type="text" placeholder="name" v-model="categoryName" />
  </div>
  {{ message }}
  <button @click="addCategory">Add category</button>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      categoryName: "",
      message: "",
      newCategory: [],
    };
  },
  methods: {
    async addCategory() {
      this.newCategory = {
        Name: this.categoryName,
      };
      try {
        await axios.post("http://localhost:8090/addCategory", this.newCategory);
        this.message = "Category added";
        console.log(this.newCategory);
      } catch (err) {
        console.log(err);
        this.message = "Error with adding category, checks inputs";
      }
    },
  },
};
</script>

<style  scoped>
body {
  font-family: Arial, sans-serif;
  background: url(http://www.shukatsu-note.com/wp-content/uploads/2014/12/computer-564136_1280.jpg)
    no-repeat;
  background-size: cover;
  height: 100vh;
}

.overlay {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.7);
  transition: opacity 500ms;
  visibility: hidden;
  opacity: 0;
}
.overlay:target {
  visibility: visible;
  opacity: 1;
}

.popup {
  margin: 70px auto;
  padding: 20px;
  background: #fff;
  border-radius: 5px;
  width: 30%;
  position: relative;
  transition: all 5s ease-in-out;
}

.popup h2 {
  margin-top: 0;
  color: #333;
  font-family: Tahoma, Arial, sans-serif;
}
.popup .close {
  position: absolute;
  top: 20px;
  right: 30px;
  transition: all 200ms;
  font-size: 30px;
  font-weight: bold;
  text-decoration: none;
  color: #333;
}
.popup .close:hover {
  color: #06d85f;
}
.popup .content {
  max-height: 30%;
  overflow: auto;
}

@media screen and (max-width: 700px) {
  .box {
    width: 70%;
  }
  .popup {
    width: 70%;
  }
}
</style>