import { createApp } from "vue";
import App from "./App.vue";
import router from "./router/index";
import store from "./store/index";
import axios from "./http";
import installElementPlus from "./plugins/element";
import ElementPlus from "element-plus";
import Region from "v-region";
const app = createApp(App);
app.config.globalProperties.$axios = axios;
installElementPlus(app);
app.use(store).use(ElementPlus).use(router).use(Region).mount("#app");
