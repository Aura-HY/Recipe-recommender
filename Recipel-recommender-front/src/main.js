import App from './App.vue';
import Vant from 'vant';
import { createApp } from 'vue';
import router from './router/router.js';
import 'vant/lib/index.css';

const app = createApp(App);

app.use(Vant);
app.use(router);
app.mount('#app');