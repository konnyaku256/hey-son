<template>
  <div class="heyson-button">
    <button @click="heySon">{{ buttoName }}</button>
  </div>
</template>

<script>
import axios from "@/assets/axios.js";
export default {
  name: "HeySonButton",
  props: {
    buttoName: String,
    buttonKind: String
  },
  methods: {
    heySon: async function() {
      await axios.get("/api/v1/heyson?kind=" + this.buttonKind);
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
