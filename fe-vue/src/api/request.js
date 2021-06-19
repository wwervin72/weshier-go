import axios from "axios";
// import store from '@/store'
import router from "@/router";
import { message } from "ant-design-vue";
import store from "@/store";

// create an axios instance
const service = axios.create({
	baseURL: process.env.VUE_APP_BASE_API,
	// withCredentials: true, // send cookies when cross-domain requests
	timeout: 10000 // request timeout
});

// request interceptor
service.interceptors.request.use(
	config => {
		config.headers.Authorization = "Bearer " + store.getters.token;
		return config;
	},
	error => {
		// do something with request error
		return Promise.reject(error);
	}
);

// response interceptor
service.interceptors.response.use(
	res => {
		const { data } = res
		const msg = data.message;
		if (data.code === 20401) {
			message.destroy();
			message.error("登录过期，请重新登录！");
			store.dispatch("handleLogout");
			router.replace({
				path: "/login" // 到登录页重新获取token
			});
		} else if (msg) {
			message[res.data.code === 200 ? "success" : "warning"](msg);
		}
		return Promise[res.data.code === 200 ? 'resolve' : 'reject'](res.data)
	},
	error => {
		message.error("请求错误");
		return Promise.reject(error.response.data);
	}
);

export default service;
