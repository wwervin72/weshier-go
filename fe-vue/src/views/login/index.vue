<template>
	<div class="ws_login">
		<router-link to="/">
			<img class="logo" :src="require('../../assets/logo/weshier.png')" title="weshier"/>
		</router-link>
		<a-form-model
			class="login_form"
			ref="loginForm"
			:model="loginObj"
			:rules="loginRules"
			v-bind="layout"
		>
			<h2 class="login_title">登录</h2>
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
			<div class="login_shortcut">
				<p>快捷登录</p>
				<a-icon class="shortcut_icon" type="github" @click="githubLoginHandler"/>
			</div>
		</a-form-model>
	</div>
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
<style lang="scss" scoped>
@import "@/styles/_variables.scss";
.logo {
	position: absolute;
	width: 120px;
	left: 42px;
	top: 42px;
}
.login_form {
	width: 480px;
	margin: 150px auto;
	transform: translate(100%, 0);
	padding: 50px 20px;
	background-color: #fff;
	box-shadow: 0 0 8px rgba(#000, 0.1);
}
.login_title {
	text-align: center;
	margin-bottom: 15px;
	color: $pinkColor;
	letter-spacing: 5px;
}
.login_shortcut {
	position: relative;
	text-align: center;
	color: $grayColor;
	&::before,
	&::after {
		position: absolute;
		top: 11px;
		content: '';
		display: block;
		width: 30%;
		height: 1px;
		background-color: $borderColor;
	}
	&::before {
		left: 10%;
	}
	&::after {
		right: 10%;
	}
	p {
		margin-bottom: 10px;
	}
}
.shortcut_icon {
	font-size: 18px;
	&:hover {
		color: $pinkColor;
	}
}
/deep/ {
	.ant-form-item {
		display: flex;
	}
	.ant-form-item-label {
		flex: 0 0 80px;
	}
	.ant-form-item-control-wrapper {
		flex: 1;
		// width: 360px;
	}
}
</style>
