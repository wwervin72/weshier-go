<template>
	<a-tooltip :title="title">
		<router-link class="ws_avatar" :style="style" v-if="to" :to="to">
			<img :src="avatar">
		</router-link>
		<img v-else class="ws_avatar" :style="style" :src="avatar">
	</a-tooltip>
</template>

<script>
import { Tooltip } from 'ant-design-vue'
export default {
	name: 'WsAvatar',
	components: {
		ATooltip: Tooltip
	},
	props: {
		to: {
			type: String,
			default: ''
		},
		title: {
			type: String,
			default: process.env.VUE_APP_NAME
		},
		avatar: {
			type: String,
			default: require('../assets/header/defaultUserAvatar.png')
		},
		size: {
			type: Number,
			default: 35
		}
	},
	data(){
		return {}
	},
	computed: {
		style () {
			return {
				width: this.size + 'px',
				height: this.size + 'px'
			}
		}
	}
}
</script>
<style lang="scss" scoped>
.ws_avatar {
	display: inline-block;
	border-radius: 50%;
	margin-right: 10px;
	cursor: pointer;
	img {
		width: 100%;
		height: 100%;
	}
}
</style>
