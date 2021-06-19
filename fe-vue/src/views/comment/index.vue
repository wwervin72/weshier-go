<template>
	<div class="comment_wrap">
		<ws-comment-editor class="comment_editor" @add-comment="addComment"></ws-comment-editor>
		<ws-comment v-for="comment in comments" :key="comment.id" :comment="comment"></ws-comment>
		<a-pagination class="text_rt ws_pagination" :default-current="pageNumber" @change="pagination" :total="total" />
	</div>
</template>

<script>
import { commentPagination } from '@/api/fetch/comment'
import WsCommentEditor from '@/components/commentEditor'
import WsComment from '@/components/comment'
export default {
	name: "WsCommentPage",
	components: {
		WsCommentEditor,
		WsComment
	},
	data() {
		return {
			comments: [],
			pageNumber: 0,
			pageSize: 10,
			total: 0
		}
	},
	created() {
		this.commentPagination()
	},
	methods: {
		addComment(data) {
			console.log(data);
		},
		commentPagination() {
			commentPagination({
				pageSize: this.pageSize,
				pageNumber: this.pageNumber
			}).then(res => {
				if (res.code === 200) {
					this.total = res.data.data.total
					this.comments = res.data.data.list
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		pagination(pageNumber) {
			this.pageNumber = pageNumber - 1
			this.commentPagination()
		}
	}
}
</script>
<style lang="scss" scoped>
.comment_editor {
	margin: 10vh 0;
}
.comment_wrap {
	width: 50%;
	min-width: 720px;
	margin: 0 auto;
	padding: 50px 0;
}
</style>
