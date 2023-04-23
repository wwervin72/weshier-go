<template>
	<footer>
		<div class="bio">
			<div class="app_desc">
				<ws-avatar :avatar="logo" to="/"></ws-avatar>
				<p>wesher 是一个博客记录站点</p>
			</div>
			<div class="about_me">
				<h1>关于</h1>
				<a-tooltip title="https://github.com/wwervin72">
					<a href="https://github.com/wwervin72">
						<a-icon type="github" />
					</a>
				</a-tooltip>
				<a-tooltip title="mailto:lw_ervin@sina.cn">
					<a href="mailto:lw_ervin@sina.cn">
						<a-icon type="mail" />
					</a>
				</a-tooltip>
				<a-tooltip title="2411120974">
					<a-icon type="qq"/>
				</a-tooltip>
			</div>
		</div>
		<div class="app_info">
			<span>
				<a href="https://beian.miit.gov.cn/">渝ICP备2021011912号-1</a>&nbsp;
				&hearts;&nbsp;<a href="mailto:lw_ervin@sina.cn">ervin</a>
			</span>
			<span class="power_by">
				power by golang & vue
			</span>
		</div>
	</footer>
</template>

<script>
import { Tooltip, Icon } from 'ant-design-vue'
import wsAvatar from './avatar'
export default {
	name: "WsFooter",
	components: {
		wsAvatar,
		AIcon: Icon,
		ATooltip: Tooltip
	},
	data() {
		return {
			logo: require('@/assets/logo/shier.png')
		}
	}
}
</script>
<style lang="scss" scoped>
@import "@/styles/_variables.scss";
footer {
	padding: 50px 0;
	background-color: #161a1d;
	&, a, h1 {
		color: $grayColor;
	}
}
.bio {
	width: 70%;
	min-width: 720px;
	margin: 0 15% 15px;
	display: flex;
	justify-content: space-around;
}
.about_me {
	a, .anticon {
		font-size: 30px;
		margin-right: 10px;
	}
}
.app_info {
	text-align: center;
	vertical-align: middle;
}
.power_by {
	margin-left: 20px;
}
</style>
