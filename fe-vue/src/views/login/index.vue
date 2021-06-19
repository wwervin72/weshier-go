<template>
	<a-form-model
		ref="loginForm"
		:model="loginObj"
		:rules="loginRules"
		v-bind="layout"
	>
		<a-form-model-item has-feedback label="账号" prop="userName">
			<a-input
				v-model="loginObj.userName"
				autocomplete="off"
			/>
		</a-form-model-item>
		<a-form-model-item has-feedback label="密码" prop="passWord">
			<a-input
				v-model="loginObj.passWord"
				type="password"
				autocomplete="off"
			/>
		</a-form-model-item>
		<a-form-model-item :wrapper-col="{ span: 14, offset: 4 }">
			<a-button type="primary" @click="submit">
				登录
			</a-button>
		</a-form-model-item>
		<a-icon type="github" @click="githubLoginHandler" width="20" height="20"/>
	</a-form-model>
</template>

<script>
import { mapActions } from 'vuex'
import { FormModel as AFormModel, Input as AInput, Button as AButton, Icon } from "ant-design-vue";
import { loginJump } from '@/utils/utils'
export default {
	name: "WsHomePage",
	components: {
		AFormModel,
		AInput,
		AButton,
		AIcon: Icon,
		AFormModelItem: AFormModel.Item
	},
	data() {
		return {
			loginObj: {
				userName: 'admin',
				passWord: 'admin'
			},
			loginRules: {
				password: [
					{
						required: true,
						message: "请输入用户名",
						trigger: "change"
					}
				],
				username: [
					{ required: true, message: "请输入密码", trigger: "change" }
				]
			},
			layout: {
				labelCol: { span: 4 },
				wrapperCol: { span: 14 }
			},
			githubLoginWindow: null
		};
	},
	methods: {
		...mapActions(['login', 'githubLogin', 'handleLogin']),
		onmessage(event) {
			var origin = event.origin || event.originalEvent.origin;
			if (origin === window.location.origin) {
				let data = event.data
				data = data && data.trim && data.trim()
				if (!data) return
				try {
					data = JSON.parse(event.data)
				} catch(e) {
					if (process.env.NODE_ENV === 'development') {
						console.log(e)
					}
				}
				if (data.close) {
					this.githubLoginWindow && this.githubLoginWindow.close()
				}
				if (data.loginSuccess && data.userInfo) {
					this.handleLogin(data.userInfo)
					loginJump.call(this)
				}
			}
		},
		submit() {
			this.$refs.loginForm.validate(valid => {
				if (valid) {
					this.login(this.loginObj).then(() => {
						loginJump.call(this)
					}).catch(err => {
						if (process.env.NODE_ENV === 'development') {
							console.log(err)
						}
					})
				} else {
					return false;
				}
			});
		},
		githubLoginHandler() {
			const redirectUrl = `${window.location.origin}/githubLogin`
			this.githubLoginWindow = window.open(
				`https://github.com/login/oauth/authorize?client_id=${process.env.VUE_APP_GITHUB_CLIENT_ID}&redirect_uri=${redirectUrl}`,
				'githubLogin',
				'resizable, scrollbars, status, width=500, height=300'
			)
			this.githubLoginWindow.onmessage = this.onmessage
		}
	}
};
</script>
