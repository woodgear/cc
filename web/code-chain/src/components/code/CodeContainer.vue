<script setup lang="ts">
import { getCode } from "./code";

const filename = "bundle/nginx-1.19.9/src/core/ngx_connection.c";
const code = getCode();

function onMouseDown() {
  console.log("onMouseDown");
}
function onMouseUp() {
  console.log("onMouseUp");
}
function onMouseMove() {
  console.log("onMouseMove");
}
const pos = {
  left: "0",
};
</script>
<template>
  <div :style="{ left: pos.left + 'px' }" class="main debug">
    <div
      class="drag-holder"
      @mousemove="onMouseMove"
      @mousedown="onMouseDown"
      @mouseup="onMouseUp"
    >
      <p>{{ filename }}</p>
    </div>
    <highlightjs class="code" lang="js" :code="code" />
  </div>
</template>

<style scoped>
.debug {
  border: 1px solid red;
  overflow: scroll;
  overflow-x: scroll;
}
.main {
  width: 500px;
  height: 400px;
  overflow: scroll;
  overflow-x: scroll;
  position: relative;
}
.drag-holder {
  cursor: move;
}
.code {
  border-top: 1px solid red;
}
</style>
