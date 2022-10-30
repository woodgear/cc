import 'highlight.js/styles/github.css';
import 'highlight.js/lib/common';
import hljsVuePlugin from '@highlightjs/vue-plugin'
import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";
import "./assets/main.css";
const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(hljsVuePlugin);

app.mount("#app");
