import { createApp } from 'vue';

import { library } from '@fortawesome/fontawesome-svg-core';
import * as Icons from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import VueSweetalert2 from 'vue-sweetalert2';
import { BACKEND_URL } from '@/constants/app';
import axios from 'axios';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/css/main.scss';

const app = createApp(App);

app.use(store);
app.use(router);

// Axios.
axios.defaults.baseURL = BACKEND_URL;

// SweetAlert2.
// TODO quando fazer versão com Composition API:
// When using "Vue3: Composition API" it is better not to use this wrapper.
// It is more practical to call sweetalert2 directly.
const options = {
  backdrop: true,
};
app.use(VueSweetalert2, options);

// FontAwesome.
app.component('FontAwesomeIcon', FontAwesomeIcon);
const iconList = Object
  .keys(Icons)
  .filter((key) => key !== 'fas' && key !== 'prefix')
  .map((icon) => Icons[icon]);
library.add(...iconList);

app.mount('#app');
