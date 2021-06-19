<template>
	<a-spin :spinning="loading">
		<div class="grid-wrap">
			<div class="grid" ref="gridDom">
				<a href="#" class="grid__item" v-for="(article, index) in articles" :key="article.id" :style="{
					width: gridItemWidth + 'px'
				}">
					<div class="grid__item-bg"></div>
					<router-link :to="`/a/${article.id}`">
						<h3 class="grid__item-title">{{article.title}}</h3>
					</router-link>
					<div class="grid__item-wrap" @click.prevent="openItem(index, $event)">
						<img class="grid__item-img" :src="`${assetsBaseUrl + article.thumbnail}`" :title="`${article.title}`">
					</div>
					<p class="grid__item-content">{{article.abstract}}</p>
					<div class="grid__item-base">
						<div class="grid__item-info">
							<router-link class="title" :to="`/u/${article.authorId}`">
								{{article.author.nickName}}
							</router-link>
							<a href="javascript:;">发布于 {{article.createdAt}}</a>
							<router-link class="heart_num ws_iconfont ws_heart" :to="`/a/${article.id}#heart_row`">
								{{0}}
							</router-link>
							<router-link class="heart_num ws_iconfont ws_comment" :to="`/a/${article.id}#comment`">
								{{0}}
							</router-link>
						</div>
						<h4 class="grid__item-tag">
							<router-link v-for="tagEntity in article.tagsEntity" :key="tagEntity.id" class="tag_item" :to="`/u/${article.authorId}/t/${tagEntity.tag.id}`">
								<a-tag color="pink">
									{{tagEntity.tag.name}}
								</a-tag>
							</router-link>
						</h4>
					</div>
				</a>
			</div>
		</div>
		<div class="content">
			<div class="content__item" ref="content">
				<h3 class="content__item-title" ref="title">
					<span v-for="(item, index) in titleCharm" :key="item + index">{{item}}</span>
				</h3>
				<div class="content__item-intro">
					<img class="content__item-img" :src="assetsBaseUrl + (currentArticle ? currentArticle.thumbnail : '')" :title="currentArticle && currentArticle.title"/>
				</div>
				<div class="content__item-text">
					<div class="markdown-body" v-html="articleContent"></div>
				</div>
			</div>
			<button class="content__close" @click="closeItem">Close</button>
		</div>
	</a-spin>
</template>
<script>
import { Tag, Spin } from 'ant-design-vue';
import imagesLoaded from 'imagesloaded';
import { Grid, Content } from "./index.js";
import Masonry from 'masonry-layout';
import marked from "marked";
import highlight from 'highlight.js';
import DOMPurify from 'dompurify';

export default {
	name: "WsArticle",
	components: {
		ATag: Tag,
		ASpin: Spin
	},
	props: {
		articles: {
			type: Array,
			default: () => []
		}
	},
	data() {
		return {
			assetsBaseUrl: process.env.VUE_APP_PUBLIC_BASE_URL,
			grid: null,
			gridItemWidth: 460,
			itemIndex: 0,
			content: null,
			loading: true
		};
	},
	watch: {
		articles: {
			deep: true,
			handler() {
				this.$nextTick(() => {
					this.init()
				})
			}
		}
	},
	computed: {
		articleContent() {
			const d = this.currentArticle ? marked(this.currentArticle.content || "", {
				langPrefix: 'lang-',
				sanitizer: function (str) {
					const res = DOMPurify.sanitize(str, {RETURN_DOM: true})
					return res
				},
				highlight: function(code, language) {
					const validLanguage = highlight.getLanguage(language) ? language : 'plaintext';
					return highlight.highlight(validLanguage, code).value;
				},
			}) : ''
			return d
		},
		titleCharm() {
			return this.currentArticle ? this.currentArticle.title.split(/\s*/) : []
		},
		currentItem() {
			return this.grid ? this.grid.items[this.itemIndex] : null
		},
		currentArticle() {
			return this.articles[this.itemIndex]
		}
	},
	mounted() {
		this.content = new Content(this.$refs.content)
	},
	methods: {
		init() {
			imagesLoaded(this.$refs.gridDom.querySelectorAll('.grid__item-img'), () => {
				this.loading = false;
				this.grid = new Grid(this.$refs.gridDom);
				this.grid.initEvents(this.content)
				new Masonry(this.grid.DOM.el, {
					itemSelector: '.grid__item',
					columnWidth: this.gridItemWidth,
					gutter: 50,
					fitWidth: true
				});
			});
		},
		openItem(index) {
			this.itemIndex = index
			this.$nextTick(() => {
				this.content.DOM.titleLetters = this.$refs.title.querySelectorAll("span");
				this.content.titleLettersTotal = this.titleCharm ? this.titleCharm.length : 0
				this.grid.openItem(this.currentItem, this.content);
			})
		},
		closeItem() {
			this.grid.closeItem(this.content);
		}
	}
};
</script>
<style lang="scss" scoped>
$ColorText: #4800d4;
$ColorBg: #e8e8e8;
$ColorLink: #ec1752;
$ColorLinkHover: #eb1851;
$ColorInfo: #272526;
$gridItemBg: #f1f1f1;
.icon {
	display: block;
	width: 1.5em;
	height: 1.5em;
	margin: 0 auto;
	fill: currentColor;
}
main {
	position: relative;
}
.frame {
	position: relative;
	padding: 1.5rem 2.5rem;
}
.frame a:hover {
	color: #4c33f7;
}
/* Header */
.codrops-header {
	position: relative;
	z-index: 100;
	text-align: center;
	font-size: 1rem;
}

.codrops-header__title {
	font-size: 1rem;
	font-weight: normal;
	margin: 0;
	padding: 0;
}

.info {
	margin: 0 0 0 1.25em;
	color: $ColorInfo;
}

.github {
	display: block;
	margin: 0.15em 0.15em 0.15em 0.5em;
	padding: 0.25em;
}

.title {
	text-align: center;
}

.title__name,
.title__sub {
	font-weight: normal;
	margin: 0;
	font-size: 1rem;
}

.title__sub {
	position: relative;
	margin: 0;
}

/* Top Navigation Style */
.codrops-links {
	position: relative;
	display: flex;
	justify-content: center;
	text-align: center;
	white-space: nowrap;
}

.codrops-icon {
	display: inline-block;
	margin: 0.15em;
	padding: 0.25em;
}

.grid-wrap {
	position: relative;
	margin: 0 auto;
}

.grid-wrap--hidden {
	height: 0;
	overflow: hidden;
}
/deep/ {
	.full_grid {
		position: fixed;
		width: 100vw;
		height: 100vh;
		top: 0;
		left: 0;
		z-index: 1000;
	}
}

.grid {
	margin: 0 auto;
	position: relative;
	display: none;
}

.grid {
	display: block;
}

.grid__item {
	position: relative;
	padding: 15px 15px 8px 15px;
	margin-bottom: 5rem;
}

.grid__item-wrap {
	position: relative;
	max-width: 100%;
}

.grid__item-bg {
	position: absolute;
	top: 0;
	left: 0;
	width: 100%;
	height: 100%;
	background: $gridItemBg;
}

.grid__item-img {
	pointer-events: none;
	position: relative;
	max-width: 100%;
	margin: 0 auto;
	display: block;
	outline: 1px solid transparent;
}

.grid__item:nth-child(even) .grid__item-img {
	transform: rotate3d(0, 0, 1, 3deg);
}

.grid__item:nth-child(odd) .grid__item-img {
	transform: rotate3d(0, 0, 1, -3deg);
}

.grid__item-title {
	padding-bottom: 10px;
	margin: 5px 0 0 10px;
	transform: translate(0, 0);
}
.grid__item-base {
	display: flex;
	align-items: center;
	justify-content: space-between;
}
.grid__item-info,
.grid__item-tag {
	display: inline-block;
	font-size: 12px;
	transform: translate(0, 0);
	vertical-align: middle;
	padding-top: 10px;
}
.grid__item-content {
	width: 100%;
	padding-top: 20px;
	transform: translate(0, 0);
}
.grid__item-tag {
	.tag_item {
		display: inline-block;
		/deep/ {
			.ant-tag {
				margin-right: 4px;
			}
		}
		&:last-child /deep/ .ant-tag {
			margin-right: 0;
		}
	}
	/deep/ {
		.ant-tag {
			padding: 3px;
		}
	}
}

.content {
	margin: 0 auto;
	grid-template-columns: 100%;
	grid-template-rows: 100%;
	align-items: center;
	pointer-events: none;
	position: fixed;
	top: 0;
	left: 0;
	width: 100vw;
	z-index: 1000;
}
.content__item {
	padding: 10rem 5vw 10rem;
	grid-area: 1 / 1 / 1 / 1;
}
.content__item {
	height: 0;
	opacity: 0;
	overflow: hidden;
	padding: 0;
	pointer-events: none;
}

.content__item--current {
	height: auto;
	opacity: 1;
	padding: 3rem 5vw 10rem;
	pointer-events: auto;
}

.content__item-intro {
	position: relative;
	grid-template-columns: 100%;
	grid-template-rows: 100%;
	align-items: center;
}

.content__item-img {
	position: relative;
	height: auto;
	max-width: 60vw;
	display: block;
	margin: 20px auto;
	visibility: hidden;
}
.content__item-title {
	position: relative;
	font-size: 50px;
	line-height: 1;
	letter-spacing: 3px;
	text-align: center;
}

.content__item-title > span {
	white-space: pre;
	display: inline-block;
}

.content__item-subtitle {
	text-align: center;
	font-size: 1.25rem;
	font-weight: normal;
	margin: 3rem auto;
}

.content__item-text {
	text-align: justify;
	max-width: 800px;
	margin: 0 auto;
}

.content__item-text p {
	margin: 0;
}

.content__close {
	position: absolute;
	top: 0.75rem;
	left: 50%;
	z-index: 10000;
	transform: translateX(-50%);
	background: none;
	border: 0;
	margin: 0;
	padding: 0;
	cursor: pointer;
	color: $ColorText;
}

.content__close:focus {
	outline: none;
}

.content__item--current ~ .content__close {
	pointer-events: auto;
}

.content__indicator {
	position: absolute;
	top: calc(100vh - 6rem);
	left: calc(50% - 0.75rem);
	display: none;
}

.content__item-title > span,
.content__item-subtitle,
.content__item-text,
.content__close,
.content__indicator {
	opacity: 0;
}
// @media screen and (min-width: 55em) {
// 	.frame {
// 		display: grid;
// 		align-items: start;
// 		justify-items: start;
// 		grid-template-columns: 40% 60%;
// 		grid-template-areas: "title header";
// 	}
// 	.codrops-header {
// 		grid-area: header;
// 		justify-self: end;
// 		display: flex;
// 		flex-direction: row;
// 		align-items: flex-start;
// 		align-items: center;
// 		text-align: left;
// 	}
// 	.codrops-links {
// 		margin: 0 0 0 1.5rem;
// 	}
// 	.title {
// 		grid-area: title;
// 		display: flex;
// 		text-align: left;
// 	}
// 	.title__sub {
// 		position: relative;
// 		padding: 0 0 0 5rem;
// 		margin: 0 0 0 1rem;
// 	}
// 	.title__sub::before {
// 		content: "";
// 		height: 1px;
// 		width: 4rem;
// 		background: currentColor;
// 		position: absolute;
// 		top: 0.65rem;
// 		left: 0;
// 	}
// 	.title__sub-works {
// 		display: block;
// 	}
// 	.grid__item-tag {
// 		right: -2.45rem;
// 	}
// 	.grid__item-title {
// 		margin-left: -0.25rem;
// 	}
// 	.content__item-subtitle {
// 		font-size: 3vw;
// 	}
// 	.content__item-text {
// 		column-count: 2;
// 		column-gap: 20px;
// 	}
// 	.content__item-img {
// 		max-width: none;
// 		height: calc(100vh - 6rem);
// 	}
// 	.content__item-subtitle {
// 		max-width: 50%;
// 	}
// 	.content__indicator {
// 		display: block;
// 	}
// }
</style>
