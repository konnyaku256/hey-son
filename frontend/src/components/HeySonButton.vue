<template>
  <div class="heyson-button">
    <button @click="heySon">{{ buttoName }}</button>
  </div>
</template>

<script>
import axios from "axios";

axios.defaults.baseURL = "http://192.168.0.16:3000";
axios.defaults.headers.post["Content-Type"] = "application/json;charset=utf-8";
axios.defaults.headers.post["Access-Control-Allow-Origin"] = "*";

export default {
  name: "HeySonButton",
  props: {
    buttoName: String
  },
  methods: {
    heySon: async function() {
      await axios.get("/api/v1/heyson");
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
