<template>
	<div class="ws_articles">
		<div class="tag_list" ref="tags" v-if="tagList.length">
			<strong>标签：</strong>
			<router-link class="tag" v-for="tag in tagList" :key="tag.value" :to="`/t/${tag.value}`">
				<a-tag color="pink">{{tag.label}}</a-tag>
			</router-link>
		</div>
		<ws-articles :articles="articles" ref="articles"></ws-articles>
		<a-pagination class="text_rt ws_pagination" :default-current="pageNumber" show-quick-jumper @change="pagination" :total="total" />
		<back-top></back-top>
	</div>
</template>

<script>
import { Pagination, Tag, BackTop } from 'ant-design-vue';
import { articlePagination } from '@/api/fetch/article';
import { fetchTagList } from '@/api/fetch/tag';
import WsArticles from '@/components/articles/index.vue';
import { TweenMax, Back } from 'gsap';

export default {
	name: "WsArticlesPage",
	components: {
		WsArticles,
		ATag: Tag,
		BackTop,
		APagination: Pagination
	},
	data() {
		return {
			total: 100,
			pageNumber: 0,
			pageSize: 0,
			articles: [],
			tagList: [],
			tagLetters: null
		}
	},
	created() {
		this.articlePagination()
		this.fetchTagList()
	},
	methods: {
		fetchTagList() {
			fetchTagList().then(res => {
				if (res.code === 200) {
					this.tagList = res.data
					this.$nextTick(() => {
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
					})
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		articlePagination () {
			articlePagination({
				pageSize: this.pageSize,
				pageNumber: this.pageNumber
			}).then(res => {
				if (res.code === 200) {
					this.total = res.data.data.total
					this.articles = res.data.data.list
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		pagination(pageNumber) {
			this.pageNumber = pageNumber - 1
			this.articlePagination()
		}
	}
}
</script>
<style lang="scss" scoped>
.ws_articles {
	width: 90%;
	min-width: 768px;
	margin: 0 auto;
}
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
