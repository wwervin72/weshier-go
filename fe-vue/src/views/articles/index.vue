<template>
	<div class="ws_articles">
		<tags-row :tagList="tagList" ref="tagRow"></tags-row>
		<ws-articles :articles="articles" ref="articles"></ws-articles>
		<pagination size="small" class="text_rt ws_pagination" :default-current="pageNumber" show-quick-jumper @change="pagination" :total="total" />
	</div>
</template>

<script>
import { Pagination } from 'ant-design-vue';
import { articlePagination } from '@/api/fetch/article';
import { fetchTagList } from '@/api/fetch/tag';
import WsArticles from '@/components/articles/index.vue';
import tagsRow from '@/components/tagsRow.vue';

export default {
	name: "WsArticlesPage",
	components: {
		WsArticles,
		tagsRow,
		Pagination
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
						this.$refs.tagRow.initAni()
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
	padding: 25px 0;
	margin: 0 auto;
}

.backtop {
	right: 50px;
}
</style>
