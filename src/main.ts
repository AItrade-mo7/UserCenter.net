import '@/assets/js/AITrade.net';
import 'normalize.css';
import '@/assets/css/global.less';

import { registerSW } from 'virtual:pwa-register';
registerSW({
  onOfflineReady() {},
});

if (ViteConst) {
  window.ViteConst = {
    ...ViteConst,
    rmAgin: 'mo777',
  };
}

import { createApp } from 'vue';

import App from '@/lib/router/App.vue';
import { router } from '@/lib/router';

const app = createApp(App);

app.use(router);

app.mount('#VueApp');
