<template>
	<div id="app">
		<WsBgAni></WsBgAni>
		<router-view></router-view>
	</div>
</template>

<script>
import { mapGetters } from "vuex";
import figlet from 'figlet';
import Standard from 'figlet/importable-fonts/Standard.js';
import WsBgAni from '@/components/bgAni.vue';
export default {
	name: "App",
	components: {
		WsBgAni
	},
	data() {
		return {};
	},
	computed: {
		...mapGetters(["token"])
	},
	created() {
		this.createFigLet()
		if (this.token) {
			this.$store.dispatch('fetchUserInfo')
		}
	},
	methods: {
		createFigLet () {
			figlet.parseFont('Standard', Standard);
			figlet.text('hello weshier !', {
				font: 'Standard',
				horizontalLayout: 'default',
				verticalLayout: 'default',
				width: 100,
				whitespaceBreak: true
			}, function(err, data) {
				if (!err) {
					console.log(data);
				}
			});
		}
	}
};
</script>

<style lang="scss" scoped>
#app {
	height: 100%;
}
</style>
