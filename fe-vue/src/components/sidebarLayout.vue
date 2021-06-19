<template>
	<div class="ws_sidebar_layout">
		<ws-sidebar class="ws_sidebar"></ws-sidebar>
		<router-view class="ws_sidebar_cnt"></router-view>
	</div>
</template>

<script>
import wsSidebar from './sidebar'
export default {
	name: "WsSidebarLayout",
	components: {
		wsSidebar
	},
	data() {
		return {}
	}
}
</script>
<style lang="scss" scoped>
.ws_sidebar_layout {
	display: flex;
	height: 100%;
}
.ws_sidebar {
	flex: 0 0 200px;
}
.ws_sidebar_cnt {
	flex: 1;
}
</style>
