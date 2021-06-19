<template>
	<div class="ws_articles">
		<ws-articles :articles="articles" ref="articles"></ws-articles>
		<a-pagination :default-current="pageNumber" show-quick-jumper @change="pagination" :total="total" />
	</div>
</template>

<script>
import { Pagination } from 'ant-design-vue'
import { articlePagination } from '../../api/fetch/article'
import WsArticles from '@/components/articles/index.vue'

export default {
	name: "WsUserTagArticlesPage",
	components: {
		WsArticles,
		APagination: Pagination
	},
	props: {
		userId: String,
		tagId: String
	},
	data() {
		return {
			total: 100,
			pageNumber: 0,
			pageSize: 0,
			articles: []
		}
	},
	created() {
		this.articlePagination()
	},
	methods: {
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
</style>
