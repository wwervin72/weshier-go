<template>
	<div class="tag_list" ref="tags" v-if="tagList.length">
		<strong>标签：</strong>
		<router-link class="tag" v-for="tag in tagList" :key="tag.value" :to="`/tag/${tag.value}`">
			<a-tag color="pink">{{tag.label}}</a-tag>
		</router-link>
	</div>
</template>

<script>
import { TweenMax, Back } from 'gsap';
import { Tag } from 'ant-design-vue';
export default {
	name: "WsTags",
	components: {
		ATag: Tag
	},
	props: {
		ATag: Tag,
		tagList: {
			type: Array,
			default: () => []
		}
	},
	data() {
		return {
			tagLetters: null
		}
	},
	methods: {
		initAni() {
			if (this.$refs.tags) {
				this.tagLetters = this.$refs.tags.querySelectorAll('a.tag')
				this.tagLetters.forEach((letter, index) => {
					TweenMax.to(letter, 0.5, {
						ease: Back.easeOut,
						delay: index * 0.05,
						startAt: { y: "-250%", x: "-100%", opacity: 0 },
						y: "0%",
						x: "0%",
						opacity: 1
					});
				});
			}
		}
	}
}
</script>

<style lang="scss" scoped>
.tag_list {
	text-align: center;
	padding: 40px 0;
	strong {
		font-size: 15px;
		display: inline-block;
		vertical-align: middle;
	}
}
.tag {
	display: inline-block;
	/deep/ {
		.ant-tag {
			cursor: pointer;
			&:hover {
				text-decoration: underline;
			}
		}
	}
}
</style>
