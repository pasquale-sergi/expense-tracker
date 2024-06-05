<template>
  <hr class="vertical" />
  <div id="new">
    <h2>Add new</h2>
    <div class="inputs">
      <div class="inputitem">
        <p style="width: 9rem">Entry type:</p>
        <select id="itemType">
          <option value="0">Expense</option>
          <option value="1">Income</option>
        </select>
      </div>
      <div class="inputitem">
        <p style="width: 9rem">Name:</p>
        <input id="name" type="text" value="" placeholder="name" />
      </div>
      <div class="inputitem">
        <p style="width: 9rem">Amount:</p>
        <input
          value="0"
          id="amount"
          name="mod"
          type="number"
          placeholder="amount"
        />
      </div>
    </div>
    <button onclick="addItem()">Add Expense</button>
  </div>
</template>

<script scoped>
</script>