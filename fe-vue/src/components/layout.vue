<template>
	<div class="ws_layout" :class="[showFooter ? 'footer_layout' : '']">
		<ws-header></ws-header>
		<div class="ws_main" ref="main">
			<keep-alive v-if="$route.meta.keepAlive">
				<router-view class="ws_body"></router-view>
			</keep-alive>
			<router-view v-else class="ws_body"></router-view>
			<ws-footer v-show="showFooter" class="ws_footer"></ws-footer>
		</div>
		<side-tools :tools="sideTools" :back-top-target="() => $refs.main"></side-tools>
	</div>
</template>

<script>
import wsHeader from './header'
import wsFooter from './footer'
import sideTools from './sideTools'
export default {
	name: "WsLayout",
	components: {
		wsHeader,
		wsFooter,
		sideTools
	},
	computed: {
		showFooter() {
			return this.$route.meta.footer !== false
		},
		sideTools() {
			return this.$route.meta.sideTools || []
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
	.ws_main {
		height: 100%;
		overflow: hidden;
		overflow-y: auto;
	}
	.ws_body {
		padding-top: 50px;
	}
}
</style>
