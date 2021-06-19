<template>
	<div class="ws_menu">
		<svg class="hidden">
			<symbol id="icon-arrow" viewBox="0 0 24 24">
				<title>arrow</title>
				<polygon
					points="6.3,12.8 20.9,12.8 20.9,11.2 6.3,11.2 10.2,7.2 9,6 3.1,12 9,18 10.2,16.8 "
				/>
			</symbol>
			<symbol id="icon-drop" viewBox="0 0 24 24">
				<title>drop</title>
				<path
					d="M12,21c-3.6,0-6.6-3-6.6-6.6C5.4,11,10.8,4,11.4,3.2C11.6,3.1,11.8,3,12,3s0.4,0.1,0.6,0.3c0.6,0.8,6.1,7.8,6.1,11.2C18.6,18.1,15.6,21,12,21zM12,4.8c-1.8,2.4-5.2,7.4-5.2,9.6c0,2.9,2.3,5.2,5.2,5.2s5.2-2.3,5.2-5.2C17.2,12.2,13.8,7.3,12,4.8z"
				/>
				<path
					d="M12,18.2c-0.4,0-0.7-0.3-0.7-0.7s0.3-0.7,0.7-0.7c1.3,0,2.4-1.1,2.4-2.4c0-0.4,0.3-0.7,0.7-0.7c0.4,0,0.7,0.3,0.7,0.7C15.8,16.5,14.1,18.2,12,18.2z"
				/>
			</symbol>
			<svg id="icon-github" viewBox="0 0 33 33">
				<title>github</title>
				<path
					d="M16.608.455C7.614.455.32 7.748.32 16.745c0 7.197 4.667 13.302 11.14 15.456.815.15 1.112-.353 1.112-.785 0-.386-.014-1.411-.022-2.77-4.531.984-5.487-2.184-5.487-2.184-.741-1.882-1.809-2.383-1.809-2.383-1.479-1.01.112-.99.112-.99 1.635.115 2.495 1.679 2.495 1.679 1.453 2.489 3.813 1.77 4.741 1.353.148-1.052.569-1.77 1.034-2.177-3.617-.411-7.42-1.809-7.42-8.051 0-1.778.635-3.233 1.677-4.371-.168-.412-.727-2.069.16-4.311 0 0 1.367-.438 4.479 1.67a15.602 15.602 0 0 1 4.078-.549 15.62 15.62 0 0 1 4.078.549c3.11-2.108 4.475-1.67 4.475-1.67.889 2.242.33 3.899.163 4.311C26.37 12.66 27 14.115 27 15.893c0 6.258-3.809 7.635-7.437 8.038.584.503 1.105 1.497 1.105 3.017 0 2.177-.02 3.934-.02 4.468 0 .436.294.943 1.12.784 6.468-2.159 11.131-8.26 11.131-15.455 0-8.997-7.294-16.29-16.291-16.29"
				></path>
			</svg>
			<symbol id="icon-menu" viewBox="0 0 119 25">
				<title>menu</title>
				<path
					d="M36,8 L36,0 L100,0 L100,8 L36,8 Z M0,8 L0,0 L24,0 L24,8 L0,8 Z M0,28 L0,20 L71,20 L71,28 L0,28 Z"
				/>
			</symbol>
			<symbol id="icon-close" viewBox="0 0 24 24">
				<title>close</title>
				<path
					d="M24 1.485L22.515 0 12 10.515 1.485 0 0 1.485 10.515 12 0 22.515 1.485 24 12 13.484 22.515 24 24 22.515 13.484 12z"
				/>
			</symbol>
		</svg>
		<div class="frame">
			<header class="codrops-header"></header>
			<p class="info">哈咯！欢迎来到<strong>ervinli</strong>的个人小站</p>
			<a
				class="github"
				href="https://github.com/wwervin72"
				title="哈咯, 我是 ervinli"
				target="_blank"
				><svg class="icon icon--github">
					<use xlink:href="#icon-github"></use></svg
			></a>
		</div>
		<div class="content">
			<div class="background">
				<div class="background__copy" v-for="item in 4" :key="item"></div>
			</div>
			<h2 class="content__title">Weshier</h2>
			<p class="content__tagline">一个博客记录网站</p>
		</div>
		<nav class="menu">
			<div class="menu__item menu__item--1" data-direction="bt">
				<div class="menu__item-inner">
					<div class="mainmenu">
						<a href="#" class="mainmenu__item">美滋滋</a>
						<a href="/articles" class="mainmenu__item">疯狂打码</a>
						<a href="#" class="mainmenu__item">小憩一下</a>
						<a href="#" class="mainmenu__item">咦！这是啥</a>
					</div>
					<p class="label label--topleft">
						the episode of my life
					</p>
					<p class="label label--bottomright">
						powered by Vue & Golang
					</p>
				</div>
			</div>
			<div class="menu__item menu__item--2" data-direction="lr">
				<div class="menu__item-inner">
					<div class="menu__item-map"></div>
					<a href="https://github.com/wwervin72" target="_blank" class="menu__item-hoverlink">about me</a>
				</div>
			</div>
			<div class="menu__item menu__item--3" data-direction="bt">
				<div class="menu__item-inner">
					<div class="sidemenu">
						<a href="#" class="sidemenu__item"
							><span class="sidemenu__item-inner"
								>FE</span
							></a
						>
						<a href="#" class="sidemenu__item"
							><span class="sidemenu__item-inner"
								>Javascript</span
							></a
						>
						<a href="#" class="sidemenu__item"
							><span class="sidemenu__item-inner"
								>CSS</span
							></a
						>
						<a href="#" class="sidemenu__item"
							><span class="sidemenu__item-inner"
								>NodeJS</span
							></a
						>
						<a href="#" class="sidemenu__item"
							><span class="sidemenu__item-inner">Golang</span></a
						>
						<a href="#" class="sidemenu__item"
							><span class="sidemenu__item-inner">Others</span></a
						>
						<a href="#" class="sidemenu__item"
							><span class="sidemenu__item-inner">...</span></a
						>
					</div>
				</div>
			</div>
			<div class="menu__item menu__item--4" data-direction="rl">
				<div class="menu__item-inner">
					<p class="label label--topleft label--line">就是玩</p>
					<a href="#" class="menu__item-link"
						>彩云之南</a
					>
				</div>
			</div>
			<div class="menu__item menu__item--5" data-direction="tb">
				<div class="menu__item-inner">
					<p class="quote">
						Vue, React, NodeJS, Golang, Vite and so on
					</p>
				</div>
			</div>
			<button class="action action--menu">
				<svg class="icon icon--menu">
					<use xlink:href="#icon-menu"></use>
				</svg>
			</button>
			<button class="action action--close">
				<svg class="icon icon--close">
					<use xlink:href="#icon-close"></use>
				</svg>
			</button>
		</nav>
	</div>
</template>
<script>
import { initMenu } from "./menu";


export default {
	name: "WsMenuBox",
	data() {
		return {};
	},
	mounted() {
		initMenu();
	}
};
</script>
<style lang="scss" scoped>
@import "./style.scss";
.ws_menu {
	width: 100vw;
	height: 100vh;
	overflow: hidden;
}
.background, .background__copy {
	position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background-image: url("./1.jpg");
    background-position: 50% 0%;
    background-size: cover;
    z-index: 1;
}
.background__copy {
    background-image: url("./1.jpg");
}
</style>
