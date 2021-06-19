<template>
	<div class="ws_bg_ani">
		<img :src="require('@/assets/bgAni/shape-1.svg')" class="shape-three"/>
		<img :src="require('@/assets/bgAni/shape-11.svg')" class="shape-four"/>
		<img :src="require('@/assets/bgAni/shape-56.svg')" class="shape-five"/>
		<img :src="require('@/assets/bgAni/shape-57.svg')" class="shape-six"/>
		<img :src="require('@/assets/bgAni/shape-58.svg')" class="shape-seven"/>
		<img :src="require('@/assets/bgAni/shape-59.svg')" class="shape-eight"/>
		<img :src="require('@/assets/bgAni/shape-10.svg')" class="shape-nine"/>
		<img :src="require('@/assets/bgAni/shape-55.svg')" class="shape-ten"/>
		<img :src="require('@/assets/bgAni/shape-62.svg')" class="shape-eleven"/>
		<img :src="require('@/assets/bgAni/shape-60.svg')" class="shape-10"/>
		<img :src="require('@/assets/bgAni/shape-61.svg')" class="shape-11"/>
	</div>
</template>

<script>
export default {
	name: "WSBgAni",
	data() {
		return {};
	}
};
</script>

<style lang="scss" scoped>
.ws_bg_ani {
	position: fixed;
	width: 100vw;
	height: 100vh;
	left: 0;
	top: 0;
	overflow: hidden;
	z-index: -1;
	// background-color: rgba($color: #f00, $alpha: 0.3);
	&,
	* {
		pointer-events: none;
	}
}
img {
	border-style: none;
	max-width: 100%;
	height: auto;
	vertical-align: middle;
}
.shape-three {
	position: absolute;
	top: 21%;
	right: 22%;
	animation: animationFramesTwo 13s linear infinite;
}
.shape-four {
	position: absolute;
	top: 27%;
	left: 7%;
	animation: animationFramesFour 25s linear infinite alternate;
}
.shape-five {
	position: absolute;
	top: 50%;
	left: 50%;
	animation: animationFramesThree 35s linear infinite alternate;
}
.shape-six {
	position: absolute;
	top: 70%;
	right: 3%;
	animation: animationFramesFour 20s linear infinite alternate;
}
.shape-seven {
	position: absolute;
	bottom: 2%;
	right: 28%;
	animation: animationFramesOne 15s linear infinite;
}
.shape-eight {
	position: absolute;
	bottom: 3%;
	left: 38%;
	animation: animationFramesFour 20s linear infinite alternate;
}
.shape-nine {
	bottom: 4%;
	animation: animationFramesOne 17s linear infinite;
}
.shape-nine,
.shape-ten {
	position: absolute;
	left: 4%;
}
.shape-ten {
	top: 60%;
	animation: animationFramesOne 20s linear infinite alternate;
}
.shape-eleven {
	top: 22%;
	left: 41%;
}
.shape-10,
.shape-eleven {
	position: absolute;
	animation: animationFramesOne 15s linear infinite;
}
.shape-10 {
	top: 2%;
	right: 2%;
}
.shape-11 {
	position: absolute;
	top: 6%;
	right: 4%;
	animation: animationFramesOne 15s linear infinite;
}
@keyframes animationFramesOne {
	0% {
		transform: translate(0) rotate(0deg);
	}
	20% {
		transform: translate(73px, -1px) rotate(36deg);
	}
	40% {
		transform: translate(141px, 72px) rotate(72deg);
	}
	60% {
		transform: translate(83px, 122px) rotate(108deg);
	}
	80% {
		transform: translate(-40px, 72px) rotate(144deg);
	}
	to {
		transform: translate(0) rotate(0deg);
	}
}
@keyframes animationFramesTwo {
	0% {
		transform: translate(0) rotate(0deg) scale(1);
	}
	20% {
		transform: translate(73px, -1px) rotate(36deg) scale(0.9);
	}
	40% {
		transform: translate(141px, 72px) rotate(72deg) scale(1);
	}
	60% {
		transform: translate(83px, 122px) rotate(108deg) scale(1.2);
	}
	80% {
		transform: translate(-40px, 72px) rotate(144deg) scale(1.1);
	}
	to {
		transform: translate(0) rotate(0deg) scale(1);
	}
}
@keyframes animationFramesThree {
	0% {
		transform: translate(165px, -179px);
	}
	to {
		transform: translate(-346px, 617px);
	}
}
@keyframes animationFramesFour {
	0% {
		transform: translate(-300px, 151px) rotate(0deg);
	}
	to {
		transform: translate(251px, -200px) rotate(180deg);
	}
}
@keyframes animationFramesFive {
	0% {
		transform: translate(61px, -99px) rotate(0deg);
	}
	21% {
		transform: translate(4px, -190px) rotate(38deg);
	}
	41% {
		transform: translate(-139px, -200px) rotate(74deg);
	}
	60% {
		transform: translate(-263px, -164px) rotate(108deg);
	}
	80% {
		transform: translate(-195px, -49px) rotate(144deg);
	}
	to {
		transform: translate(-1px) rotate(180deg);
	}
}
</style>
