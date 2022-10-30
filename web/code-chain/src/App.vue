<script setup lang="ts">
import CodeChain from "./components/codechain/CodeChain.vue";
</script>

<template>
  <header></header>
  <main>
    <CodeChain class="chain"></CodeChain>
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}
main {
  border: 1px solid black;
  display: flex;
  align-items: center;
  justify-content: center;
}
.chain {
  height: 800px;
  width: 800px;
  border: 1px solid orange;
}
</style>
