<template>
	<div class="ws_menu" :class="[hover ? 'hover' : '']" @mouseenter="enterMenu" @mouseleave="leaveMenu">
		<div class="c_menus">
			<router-link class="c_menu" to="/articles">
				<a-icon type="book" />
			</router-link>
			<router-link class="c_menu" to="/comment">
				<a-icon type="edit" />
			</router-link>
			<router-link class="c_menu" to="/login">
				<a-icon type="login" />
			</router-link>
		</div>
		<div class="trigger">
			<a-icon type="menu" />
		</div>
	</div>
</template>
<script>
import { Icon } from "ant-design-vue";
export default {
	name: "WsMenu",
	components: {
		AIcon: Icon
	},
	data() {
		return {
			hover: false
		};
	},
	methods: {
		enterMenu() {
			this.hover = true
		},
		leaveMenu() {
			this.hover = false
		}
	}
};
</script>
<style lang="scss" scoped>
@import "@/styles/_function.scss";
@import "@/styles/_variables.scss";
$distance: 180px;
$childMenuLens: 3;
$menuSize: 40px;
$bg: $pinkColor;
.ws_menu {
	position: fixed;
	width: $menuSize;
	height: $menuSize;
	top: 30px;
	right: 30px;
	color: #fff;
	transition: all 1s ease;
	transform: rotate(180deg);
}
.trigger {
	display: flex;
	position: absolute;
	border-radius: 50%;
	width: $menuSize;
	height: $menuSize;
	justify-content: center;
	align-items: center;
	background-color: $bg;
	color: #fff;
	transition: all .5s ease-in-out;
	transform: scale(.85);
}
.c_menu {
	display: flex;
	position: absolute;
	width: $menuSize;
	height: $menuSize;
	font-size: 20px;
	justify-content: center;
	align-items: center;
	background: $bg;
	cursor: pointer;
	border-radius: 50%;
	color: #fff;
	transition: all ease-in-out .5s;
}
.c_menu {
	opacity: 0;
	transition: all ease-in-out .5s;
}
.c_menus {
	position: absolute;
	width: $distance + $menuSize * 2;
	height: $distance + $menuSize * 2;
	right: -$distance / 2 - $menuSize / 2;
	top: -$distance / 2 - $menuSize / 2;
	transform-origin: center center;
	transition: all ease-in-out .3s;
}
@for $i from 0 through $childMenuLens {
	.c_menu:nth-child(#{$i + 1}) {
		top: sin(0deg + 90 / $childMenuLens * $i) * $distance;
		right: cos(0deg + 90 / $childMenuLens * $i) * $distance;
		transform: translate(-$distance / 2 + $menuSize / 2, $distance / 2);
	}
	.hover {
		.c_menus {
			transform: rotate(-190deg);
		}
		.c_menu {
			opacity: 1;
		}
		.trigger {
			transform: scale(1);
		}
	}
}
</style>
