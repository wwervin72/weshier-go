<template>
	<div class="ws_article">
		<a-skeleton v-if="loadingArticle" active />
		<template v-else>
			<h1 class="ws_title">{{article.title}}</h1>
			<div class="head_info" v-if="article.author">
				<ws-avatar :to="`/u/${article.authorId}`" :avatar="article.author.avatar || undefined"></ws-avatar>
				<div class="info_right">
					<router-link :to="`/u/${article.author.id}`">
						<span>{{article.author.nickName}}</span>
						<a-tag class="author" color="pink">作者</a-tag>
					</router-link>
					<p>
						{{article.createdAt}}
					</p>
				</div>
			</div>
			<div class="markdown-body" v-html="content"></div>
			<div class="article_info">
				<div class="info_lf">
					<a-icon
						type="heart"
						:theme="article.hearted ? 'filled' : 'outlined'"
						@click="heartArticle"
					/>
					<span>{{article.heartCount}}</span>
				</div>
				<div class="info_rt">
					<a-icon type="delete" @click="delArticle"/>
					<router-link :to="'/update/' + article.id">
						<a-icon type="edit" />
					</router-link>
				</div>
			</div>
			<div class="article_info">
				<div class="info_lf">
					<a-icon type="tag" class="tag_icon" />
					<router-link v-for="tagEntity in article.tagsEntity" :key="tagEntity.id" :to="`/u/${article.authorId}/t/${tagEntity.tag.id}`">
						<a-tag color="pink">
							{{tagEntity.tag.name}}
						</a-tag>
					</router-link>
				</div>
				<div class="info_rt">
					<router-link :to="`/u/${article.authorId}/c/${article.categoryId}`">
						<a-tag color="pink">
							{{article.categoryId}}
						</a-tag>
					</router-link>
				</div>
			</div>
		</template>
		<div id="comment">
			<ws-comment-editor class="comment_editor" :article-id="article.id" @add-comment="addComment"></ws-comment-editor>
			<a-skeleton v-if="loadingComments" active />
			<template v-else>
				<ws-comment v-for="comment in comments" :key="comment.id" :comment="comment"
					:article-id="article.id" :article-author="article.author.id"></ws-comment>
			</template>
			<a-pagination class="article_pagination" :default-current="pageNumber" @change="pagination" :total="total" />
		</div>
	</div>
</template>

<script>
import { Pagination, Icon, Tag, Skeleton } from 'ant-design-vue'
import marked from "marked";
import highlight from 'highlight.js';
import DOMPurify from 'dompurify';
import 'github-markdown-css/github-markdown.css'
// import "highlight.js/styles/default.css";
// import "highlight.js/styles/atom-one-dark.css";

import { fetchArticleInfo } from "@/api/fetch/article";
import { articleCommentPagination } from "@/api/fetch/comment";
import wsAvatar from '@/components/avatar'
import wsCommentEditor from '@/components/commentEditor'
import wsComment from '@/components/comment'
export default {
	name: "WsArticlePage",
	components: {
		ASkeleton: Skeleton,
		ATag: Tag,
		AIcon: Icon,
		APagination: Pagination,
		wsAvatar,
		wsCommentEditor,
		wsComment
	},
	props: {
		articleId: {
			type: String,
			default: ""
		}
	},
	data() {
		return {
			loadingArticle: false,
			loadingComments: false,
			article: {},
			comments: [],
			pageNumber: 0,
			pageSize: 10,
			total: 0
		};
	},
	created() {
		this.fetchArticleInfo();
		this.articleCommentPagination();
		highlight.highlightAll()
	},
	computed: {
		content() {
			const d = marked(this.article.content || "", {
				langPrefix: 'lang-',
				sanitizer: function (str) {
					console.log(arguments);
					const res = DOMPurify.sanitize(str, {RETURN_DOM: true})
					console.log(res);
					return res
				},
				highlight: function(code, language) {
					const validLanguage = highlight.getLanguage(language) ? language : 'plaintext';
					return highlight.highlight(validLanguage, code).value;
				},
			})
			return d
		}
	},
	methods: {
		delArticle() {},
		fetchArticleInfo() {
			this.loadingArticle = true
			fetchArticleInfo(this.articleId)
				.then(res => {
					if (res.code === 200) {
						this.article = res.data;
					}
				})
				.catch(err => {
					if (process.env.NODE_ENV === 'development') {
						console.log(err);
					}
				})
				.finally(() => {
					this.loadingArticle = false
				});
		},
		addComment(data) {
			console.log(data);
		},
		articleCommentPagination() {
			this.loadingComments = true
			articleCommentPagination(this.articleId, {
				pageSize: this.pageSize,
				pageNumber: this.pageNumber
			})
			.then(res => {
				if (res.code === 200) {
					this.total = res.data.data.total
					this.comments = res.data.data.list
				}
			})
			.catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
			.finally(() => {
				this.loadingComments = false
			});
		},
		pagination(pageNumber) {
			this.pageNumber = pageNumber - 1
			this.commentPagination()
		},
		heartArticle() {},
	}
};
</script>
<style lang="scss" scoped>
@import "@/styles/_variables.scss";
.ws_article {
	width: 50%;
	padding: 25px 0;
	margin-left: 25%;
}
.ws_title {
	margin-bottom: 10px;
}
.head_info {
	display: flex;
	margin: 20px 0 50px 0;
	.author {
		margin-left: 5px;
	}
}
.info_right {
	flex: 1;
}
.tag_icon {
	margin-right: 5px;
}
.article_info {
	display: flex;
	padding: 15px 0;
	border-bottom: 1px solid $borderColor;
	.info_rt {
		flex: 0 0 100px;
		text-align: right;
		/deep/ {
			.anticon {
				display: none;
			}
		}
	}
	&:hover {
		/deep/ {
			.anticon {
				display: inline-block;
			}
		}
	}
	.info_lf {
		flex: 1;
	}
	/deep/ {
		.anticon {
			font-size: 16px;
			margin-left: 10px;
		}
	}
}
.comment_editor {
	margin: 50px 0 30px 0;
}
.article_pagination {
	margin-top: 50px;
	text-align: right;
}
</style>
