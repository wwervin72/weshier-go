<template>
	<div class="ws_side_tools">
		<template v-for="item in tools">
			<back-top class="ws_back_top" :key="item" v-if="item === backTop" :target="backTopTarget">
				<icon class="ws_tool back_top_icon" type="to-top" />
			</back-top>
			<popover :key="item" v-else-if="item === toc && tocHtml" placement="left">
				<template slot="content">
					<div class="ws_toc" v-html="tocHtml"></div>
				</template>
				<icon class="ws_tool" type="menu"/>
			</popover>
		</template>
	</div>
</template>

<script>
import { Popover, Icon, BackTop } from 'ant-design-vue';
import { mapGetters } from 'vuex';
import * as tools from './config';

export default {
	name: "WsSideTools",
	props: {
		tools: {
			type: Array,
			default: () => [...Object.keys(tools)]
		},
		backTopTarget: {
			type: Object,
			default: null
		}
	},
	computed: {
		...mapGetters(['tocHtml'])
	},
	components: {
		Icon,
		Popover,
		BackTop
	},
	data() {
		return {
			...tools,
		}
	},
	methods: {}
}
</script>

<style lang="scss" scoped>
@import "@/styles/_variables.scss";
.ws_tool {
	width: 40px;
	height: 40px;
	line-height: 40px;
	border: 1px solid $borderColor;
	margin-top: -1px;
	display: block;
	cursor: pointer;
}
.ws_toc {
	/deep/ {
		li {
			&:hover {
				text-decoration: underline;
			}
		}
	}
}
.ws_side_tools {
	position: fixed;
	bottom: 50px;
	right: 50px;
	background-color: #fff;
}
.ws_back_top {
	position: static;
	right: auto;
	bottom: auto;
}
.back_top_icon {
	font-size: 16px;
}
/deep/ {
	.ant-back-top-content {
		width: 100%;
		height: 100%;
	}
}
</style>
