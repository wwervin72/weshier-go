<template>
	<div class="ws_comment_editor">
		<ws-avatar class="avatar" :size="35" :title="user ? user.nickName : 'weshier'"></ws-avatar>
		<div class="editor_wrap">
			<a-textarea
				v-model="content"
				placeholder="来都来了，总得留下点什么"
				:auto-size="{ minRows: 3, maxRows: 5 }"
				@keyup.enter="enterKeyUp"
			>
			</a-textarea>
			<div class="editor_action">
				<div class="action_left">
					<a-popover>
						<template slot="title"></template>
						<template slot="content">
							<ws-emoji></ws-emoji>
						</template>
						<span>
							<a-icon type="smile" theme="outlined" />
						</span>
					</a-popover>
				</div>
				<div class="action_right">
					<a-button type="link" class="cancel" @click="cancel">取消</a-button>
					<a-button type="primary" @click="submit">提交</a-button>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import { Button, Input, Popover, Icon, message } from 'ant-design-vue'
import { mapGetters } from 'vuex'
import WsAvatar from './avatar'
import WsEmoji from './emoji'
import { LeaveComment, CommentArticle } from '@/api/fetch/comment'
export default {
	name: "WsCommentEditor",
	props: {
		articleId: Number,
		commentId: Number,
	},
	components: {
		AIcon: Icon,
		APopover: Popover,
		AButton: Button,
		ATextarea: Input.TextArea,
		WsAvatar,
		WsEmoji
	},
	data() {
		return {
			content: ''
		}
	},
	computed: {
		...mapGetters(['user']),
		submitFn() {
			return this.articleId ? CommentArticle : LeaveComment
		}
	},
	methods: {
		enterKeyUp(e) {
			console.log(e);
		},
		submit() {
			if (!this.content) {
				message.info("请输入评论内容")
				return
			}
			const data = {
				content: this.content
			}
			if (this.articleId) {
				data.articleId = this.articleId
			}
			if (this.commentId) {
				data.commentId = this.commentId
			}
			this.submitFn(data).then(res => {
				console.log(res);
				if (res.code === 200) {
					this.content = ''
					this.$emit('add-comment', res)
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		cancel() {
			this.content = ''
			this.$emit('cancel')
		}
	}
}
</script>
<style lang="scss" scoped>
@import "@/styles/_variables.scss";
.ws_comment_editor {
	display: flex;
}
.editor_wrap {
	flex: 1;
}
.cancel {
	color: $grayColor;
}
.editor_action {
	display: flex;
	margin-top: 5px;
	.action_left {
		flex: 1;
	}
	.action_right {
		flex: 0 0 150px;
		text-align: right;
	}
}
/deep/ {
	textarea.ant-input {
		resize: none;
	}
}
</style>
