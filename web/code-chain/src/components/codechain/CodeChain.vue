<script setup lang="ts">
import Code from "../code/CodeContainer.vue";
</script>
<template>
  <div class="code debug">
    <Code />
    <Code />
  </div>
</template>

<style scoped>
.debug {
  border: 1px solid red;
}
.code {
   display: block;
}
</style>
