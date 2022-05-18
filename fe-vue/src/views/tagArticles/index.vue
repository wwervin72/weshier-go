<template>
	<div class="ws_articles">
		<tags-row :tagList="tagList" ref="tagRow"></tags-row>
		<ws-articles :articles="articles" ref="articles"></ws-articles>
		<a-pagination size="small" class="text_rt ws_pagination" :default-current="pageNumber" show-quick-jumper @change="pagination" :total="total" />
	</div>
</template>

<script>
import { Pagination } from 'ant-design-vue'
import { fetchTagInfo, fetchTagArticlePagination } from '../../api/fetch/tag'
import WsArticles from '@/components/articles/index.vue'
import tagsRow from '@/components/tagsRow.vue';

export default {
	name: "WsTagArticlesPage",
	components: {
		WsArticles,
		tagsRow,
		APagination: Pagination
	},
	props: {
		tagId: String
	},
	data() {
		return {
			total: 100,
			pageNumber: 0,
			pageSize: 0,
			articles: [],
			tagList: []
		}
	},
	created() {
		this.fetchTagArticlePagination()
		this.fetchTagInfo()
	},
	methods: {
		fetchTagInfo() {
			fetchTagInfo(this.tagId).then(res => {
				this.$set(this, 'tagList', [{
					label: res.data.name,
					value: res.data.id
				}])
				this.$nextTick(() => {
					this.$refs.tagRow.initAni()
				})
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		fetchTagArticlePagination () {
			fetchTagArticlePagination(this.tagId, {
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
			this.fetchTagArticlePagination()
		}
	}
}
</script>
<style lang="scss" scoped>
.ws_articles {
	width: 90%;
	padding: 25px 0;
	min-width: 768px;
	margin: 0 auto;
}
</style>
