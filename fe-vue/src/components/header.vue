<template>
	<header class="ws_header">
		<div class="ws_header_content">
			<a href="/" class="logo"></a>
			<a href="javascript:;" class="ws_iconfont ws_caidan" id="switch_user_menu"></a>
			<div class="user_area">
				<div class="user_right">
					<router-link to="/articles">博文</router-link>
					<router-link to="/comment">留言板</router-link>
					<template v-if="user">
						<router-link v-if="isAdmin" to="/admin">管理员</router-link>
						<router-link v-if="isAdmin" to="/editor">写点啥？</router-link>
						<div class="user_dropdown">
							<template v-if="user.authId === 0">
								<router-link :to="`/u/${user.id}`">
									<img :src="user.avatar || require('../assets/header/defaultUserAvatar.png')" class="avatar" :title="user.nickName">
								</router-link>
								<ul class="dropdown">
									<li>
										<router-link :to="`/u/${user.id}`" class="ws_iconfont ws_gerenzhuye">
										</router-link>
									</li>
									<li>
										<router-link to="/changepwd" class="ws_iconfont">
										修改密码
										</router-link>
									</li>
									<li>
										<router-link to="/setting" class="ws_iconfont ws_setting"></router-link>
									</li>
									<li>
										<a href="javascript:;" @click="logoutHandler" class="ws_iconfont ws_tuichu"></a>
									</li>
								</ul>
							</template>
							<template v-else>
								<img :src="user.avatar || require('../assets/header/defaultUserAvatar.png')" class="avatar" :title="user.nickName">
								<ul class="dropdown">
									<li><a href="javascript:;" @click="logoutHandler" class="ws_iconfont ws_tuichu"></a></li>
								</ul>
							</template>
						</div>
					</template>
					<template v-else>
						<a href="javascript:;" @click.prevent="goLogin">登录</a>
					</template>
				</div>
			</div>
		</div>
	</header>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import { roles } from '@/utils/variables.js';
import { goLogin } from '@/utils/utils.js';

export default {
	name: "WsHeader",
	data() {
		return {}
	},
	computed: {
		...mapGetters(['user']),
		isAdmin() {
			return this.user.role === roles.admin
		}
	},
	methods: {
		goLogin,
		...mapActions(['logout']),
		logoutHandler () {
			this.logout().then(() => {
				if (this.$route.meta.role && this.$route.meta.role.length) {
					goLogin.call(this)
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		}
	}
}
</script>
<style lang="scss" scoped>
@import "@/styles/_header.scss";
</style>
