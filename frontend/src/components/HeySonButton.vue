<template>
  <div class="heyson-button">
    <button @click="heySon">{{ buttoName }}</button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "HeySonButton",
  props: {
    buttoName: String
  },
  methods: {
    heySon: async function() {
      await axios.get("http://backend:3000/api/v1/heyson");
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
