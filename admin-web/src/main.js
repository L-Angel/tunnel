import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router';
import {router} from './router';
import ViewUI from 'view-design';
import {store} from "./store";


// import 'view-design/dist/styles/iview.css';
import './tunnel-theme/index.less';

Vue.config.productionTip = false



Vue.use(VueRouter);
Vue.use(ViewUI);


new Vue({
    render: h => h(App),
    router: router,
    store: store
}).$mount('#app')
