<template>
  <div class="container" style="padding: 30px" id="app">
    <h2 style="padding: 30px; text-align: center">Lista zadań</h2>

    <task-creator @ItemCreated="RefreshList"></task-creator>
    <p style="padding-bottom: 30px" />
    <task-list :items="items" @ItemDeleted="RefreshList"></task-list>
  </div>
</template>

<script>
import TaskCreator from "./components/TaskCreator";
import TaskList from "./components/TaskList";

import axios from "axios";
export default {
  name: "App",
  components: {
    TaskCreator,
    TaskList,
  },
  data: function () {
    return {
      items: [],
    };
  },
  methods: {
    async RefreshList() {
      var result = await axios.get("http://localhost:8081/items");
      this.items = result.data;
    },
  },
  created() {
    this.RefreshList();
  },
};
</script>

<style>
</style>
