<template>
  <div class="container">
    <table style="width: 100%">
      <thead>
        <tr>
          <th style="width: 85%">Zadanie</th>
          <th>Usuwanie</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item._id">
          <td>{{ item.text }}</td>
          <td>
            <button type="button" @click="DeleteItem(item._id)">Usuń</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "TaskList",
  props: ["items"],
  methods: {
    async DeleteItem(id) {
      await axios.delete("http://localhost:8081/items/" + id);
      this.$emit("ItemDeleted");
    },
  },
};
</script>

<style scoped>
table tr:nth-child(even) {
  background-color: #eeeeee;
}
table tr {
  border-bottom: 1px solid #333;
}
</style>