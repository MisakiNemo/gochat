<script setup>
import {useRoute, useRouter} from "vue-router";

const route = useRoute();
const router = useRouter();
</script>

<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<style lang="scss">
html, body {
  width: 100%;
  height: 100%;
  //background: linear-gradient(220.55deg, #DDE4FF 0%, #8DA2EE 100%);
  //background: linear-gradient(220.55deg, #4063BC 0%, #6B0013 100%);
  background-color: #e9e9e9;

  * {
    padding: 0;
    margin: 0;
  }
}

</style>
