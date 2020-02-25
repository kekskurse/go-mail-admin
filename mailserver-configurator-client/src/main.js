import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router'
import VueSweetalert2 from 'vue-sweetalert2';
import vSelect from 'vue-select'

import 'sweetalert2/dist/sweetalert2.min.css';
Vue.use(VueSweetalert2);

import 'vue-select/dist/vue-select.css';
Vue.component('v-select2', vSelect)

Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
