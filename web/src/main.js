import { createApp } from "vue";
import App from "./App.vue";
import router from "./router/index";
import store from "./store/index";
import axios from "./http";
import installElementPlus from "./plugins/element";
import ElementPlus from "element-plus";
import lazyPlugin from "vue3-lazy";
import Region from "v-region";
const app = createApp(App);
app.config.globalProperties.$axios = axios;
installElementPlus(app);
lazyPlugin.install(app, {
  loading: "loading.png",
  error: "error.png",
});
router.afterEach(() => {
  document.documentElement.scrollTop = 0;
});
app.use(store).use(ElementPlus).use(router).use(Region).mount("#app");
