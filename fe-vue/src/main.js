import Vue from "vue";
import mavonEditor from "mavon-editor";
import router from "./router";
import store from "./store";
import App from "./App.vue";
import { Pagination } from "ant-design-vue";
import { switchBrowserTabs, copySiteInfo } from "./utils/utils";

import "mavon-editor/dist/css/index.css";
import "./styles/_reset.scss";
import "./styles/_common.scss";

Vue.use(Pagination);
Vue.use(mavonEditor);
Vue.config.productionTip = false;
switchBrowserTabs();
copySiteInfo();

new Vue({
	router,
	store,
	render: h => h(App)
}).$mount("#app");
