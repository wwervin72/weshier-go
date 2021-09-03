<template>
	<a-comment class="ws_comment">
		<router-link slot="author" :to="`/u/${comment.author.id}`">
			{{comment.author.nickName}}
			<a-tag color="pink" v-if="comment.author.id === articleAuthor">作者</a-tag>
		</router-link>
		<template slot="actions">
			<span key="comment-basic-like">
				<a-tooltip title="喜欢">
					<a-icon
						type="heart"
						:style="{
							fontSize: '13px'
						}"
						:theme="comment.hearted ? 'filled' : 'outlined'"
						@click="heart"
					/>
				</a-tooltip>
				<span style="padding-left: '8px';cursor: 'auto'">{{comment.heartCount}}</span>
			</span>
			<span key="comment-basic-reply-to" @click="opReplyEditor">回复</span>
			<a-popconfirm
				title="确认删除该评论吗？"
				ok-text="删除"
				cancel-text="取消"
				@confirm="confirmDelComment"
			>
				<span key="comment-basic-reply-to" v-if="isAdmin" @click="delComment">删除</span>
			</a-popconfirm>
		</template>
		<router-link slot="avatar" :to="`/u/${comment.author.id}`">
			<ws-avatar
				:avatar="assetsBaseUrl + comment.author.avatar"
				:title="comment.author.nickName"
			/>
		</router-link>
		<span slot="datetime">发布于 {{comment.createdAt}}</span>
		<p slot="content">{{comment.content}}</p>
		<ws-comment v-for="childComment in comment.comments || []" :key="childComment.id"
			:comment="childComment" :need-editor="false" @reply="opReplyEditor" :article-author="articleAuthor" @delete-comment="delComment"></ws-comment>
		<ws-comment-editor v-if="needEditor" v-show="editorVisible" @add-comment="addComment" :metion-options="childCommentAuthors"
			@cancel="cancelComment" :articleId="articleId" :commentId="comment.id"></ws-comment-editor>
	</a-comment>
</template>

<script>
import { Comment, Icon, Tooltip, Tag, Popconfirm } from "ant-design-vue";
import { mapGetters } from 'vuex'
import WsCommentEditor from './commentEditor'
import WsAvatar from './avatar'
import { roles } from '@/utils/variables'
import { DeleteComment } from '@/api/fetch/comment'

export default {
	name: "WsComment",
	props: {
		needEditor: {
			type: Boolean,
			default: true
		},
		articleId: Number,
		articleAuthor: Number,
		comment: {
			type: Object,
			default: () => ({})
		}
	},
	computed: {
		...mapGetters(['user']),
		isAdmin() {
			return this.user && this.user.role === roles.admin
		},
		childCommentAuthors() {
			return this.comment.comments.map(el => el.author)
		}
	},
	components: {
		APopconfirm: Popconfirm,
		ATag: Tag,
		ATooltip: Tooltip,
		AIcon: Icon,
		AComment: Comment,
		WsCommentEditor,
		WsAvatar
	},
	data() {
		return {
			assetsBaseUrl: process.env.VUE_APP_PUBLIC_BASE_URL,
			editorVisible: false
		};
	},
	methods: {
		heart() {

		},
		opReplyEditor() {
			this.editorVisible = true
			this.$emit('reply')
		},
		addComment(data) {
			console.log(data);
		},
		cancelComment() {
			this.editorVisible = false
		},
		confirmDelComment() {
			DeleteComment({
				commentId: this.comment.id
			}).then(res => {
				console.log(res);
				this.$emit('delete-comment', this.comment)
			}).catch(() => {})
		},
		delComment(comment) {
			if (!Array.isArray(this.comment.comments)) {
				return
			}
			const index = this.comment.comments.findIndex(el => el.id === comment.id);
			if (index !== -1) {
				this.comment.comments.splice(index, 1)
			}
		}
	}
};
</script>
<style lang="scss" scoped>
@import '@/styles/_variables.scss';
.ws_comment {
	&::before {
		display: none;
	}
}
/deep/ {
	.ant-comment-inner {
		padding: 10px 0;
	}
	.ant-comment-nested {
		border-left: 5px solid $borderColor;
		padding-left: 10px;
	}
	.ant-comment-content-author-time {
		display: block;
		width: 100%;
	}
}
</style>
