import App from "./App.vue";
import Vant from "vant";
import { createApp } from "vue";
import router from "./router/router.js";
import "vant/lib/index.css";
import Casdoor from "casdoor-vue-sdk";

const app = createApp(App);
const config = {
  serverUrl: "http://localhost:8000",
  clientId: "4003c48c69588cb77d75",
  organizationName: "casbin",
  appName: "for_health",
  redirectPath: "/home",
};

app.use(Vant);
app.use(router);
app.mount("#app");
app.use(Casdoor, config);
