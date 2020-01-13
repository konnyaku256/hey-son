<template>
  <div class="home">
    <h2>Welcome to "Hey Son" magic</h2>
    <HeySonButton buttoName="ごはんできたよ" />
  </div>
</template>

<script>
// @ is an alias to /src
import HeySonButton from "@/components/HeySonButton.vue";

export default {
  name: "home",
  components: {
    HeySonButton
  }
};
</script>
