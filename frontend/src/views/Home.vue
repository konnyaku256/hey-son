<template>
  <div class="home">
    <h2>Welcome to "Hey Son" magic</h2>
    <HeySonButton buttoName="ごはんできたよ" buttonKind="gohan" />
    <HeySonButton buttoName="おふろ入って" buttonKind="ohuro" />
    <HeySonButton buttoName="その他の要件" buttonKind="others" />
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
