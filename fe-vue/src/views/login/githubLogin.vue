<template>
	<a-spin class="login_spin"></a-spin>
</template>

<script>
import { Spin, message } from "ant-design-vue";
import { mapActions } from "vuex";

export default {
	name: "WsGithubLogin",
	components: {
		ASpin: Spin
	},
	data() {
		return {};
	},
	created() {
		this.login();
	},
	methods: {
		...mapActions(["githubLogin"]),
		login() {
			const code = this.$route.query.code;
			if (!code) {
				message.warning("登录错误，跳转回登录界面...");
				window.location = "/login";
				return;
			}
			let loginSuccess, userInfo;
			this.githubLogin({ code })
				.then(data => {
					loginSuccess = true;
					userInfo = data;
				})
				.catch(err => {
					if (process.env.NODE_ENV === "development") {
						console.log(err);
					}
				})
				.finally(() => {
					window.postMessage(
						JSON.stringify({
							close: true,
							loginSuccess,
							userInfo
						}),
						window.location.origin
					);
				});
		}
	}
};
</script>

<style lang="scss" scoped>
.login_spin {
	display: flex;
	width: 100vw;
	height: 100vh;
	align-items: center;
	justify-content: center;
}
</style>
