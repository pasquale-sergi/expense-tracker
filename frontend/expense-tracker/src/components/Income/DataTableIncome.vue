<template>
  <table>
    <tr class="title-column">
      <th class="descr_content">Description</th>
      <th class="content">Amount</th>
    </tr>
    <tr class="columns-content" v-for="item in items" :key="item">
      <td class="descr_content">{{ item.Description }}</td>
      <td class="content">${{ item.Amount }}</td>
      <!-- Adjust columns as per your data -->
    </tr>
  </table>
</template>
  
  <script setup>
// Define props using defineProps macro
defineProps({
  items: {
    type: Array,
    required: true,
  },
});
</script>

<style scoped>
.title-column {
  margin-left: 10px;
  border: solid 2px green;
}
.descr_content {
  padding-right: 50px;
}
.content {
  padding: 10px;
}
</style>