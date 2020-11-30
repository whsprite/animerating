import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import VueConpositionApi from '@vue/composition-api'
import axios from 'axios'
Vue.prototype.$http = axios
axios.defaults.baseURL = "/api/v1/"
Vue.use(VueConpositionApi)
Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
