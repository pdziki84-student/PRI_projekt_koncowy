<template>
  <div class="container">
    <div>
      <div class="form-group">
        <input type="text" class="form-control" id="textInput" v-model="text" />
        <label for="textInput">Wpisz zadanie</label>
        <p />
      </div>
      <button type="submit" @click="CreateItem">Dodaj</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "TaskCreator",
  data: function () {
    return {
      text: "",
    };
  },
  methods: {
    async CreateItem() {
      await axios.post("http://localhost:8081/items", { text: this.text });
      this.$emit("ItemCreated");
      this.resetInput();
    },
    resetInput() {
      this.text = "";
    },
  },
};
</script>

<style scoped>
</style>