<template>
	<div class="ws_layout" :class="[showFooter ? 'footer_layout' : '']">
		<ws-header></ws-header>
		<keep-alive v-if="$route.meta.keepAlive">
			<router-view class="ws_main"></router-view>
		</keep-alive>
		<router-view v-else class="ws_main"></router-view>
		<ws-footer v-show="showFooter" class="ws_footer"></ws-footer>
	</div>
</template>

<script>
import wsHeader from './header'
import wsFooter from './footer'
export default {
	name: "WsLayout",
	components: {
		wsHeader,
		wsFooter
	},
	computed: {
		showFooter() {
			return this.$route.meta.footer !== false
		}
	},
	data() {
		return {
			notKeepAlive: [""]
		}
	}
}
</script>
<style lang="scss" scoped>
@import "@/styles/_variables.scss";
.ws_layout {
	height: 100%;
	padding-top: $headerHeight;
	&.footer_layout {
		.ws_main {
			min-height: calc(100vh - 50px);
		}
	}
}
</style>
