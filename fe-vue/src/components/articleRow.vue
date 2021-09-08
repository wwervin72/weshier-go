<template>
	<div class="article clear module_bg">
		<div v-if="article.thumbnail" class="avatar right">
			<router-link :to="`/a/${article.id}`">
				<img :src="`${assetsBaseUrl + article.thumbnail}`" :title="`${article.title}`">
			</router-link>
		</div>
		<div class="content">
			<router-link class="title" :to="`/a/${article.id}`">
				{{article.title}}
			</router-link>
			<p class="abstract">{{article.abstract}}</p>
		</div>
		<div class="base_info">
			<div class="info_lf">
				<router-link class="title" :to="`/user/${article.authorId}`">
					{{article.author.nickName}}
				</router-link>
				发布于 {{article.createdAt}}
				<router-link class="heart_num ws_iconfont ws_heart" :to="`/a/${article.id}#heart_row`">
					{{0}}
				</router-link>
				<router-link class="heart_num ws_iconfont ws_comment" :to="`/a/${article.id}#comment`">
					{{0}}
				</router-link>
			</div>
			<div class="info_rt">
				<template v-if="user && article.authorId === user.id">
					<router-link class="right edit_article ws_iconfont ws_edit" :to="`/update/${article.id}`" title="编辑">
					</router-link>
					<a class="right del_article ws_iconfont ws_shanchu" title="删除" href="javascript:;"></a>
				</template>
			</div>
		</div>
	</div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
	name: "WsArticleRow",
	props: {
		article: {
			type: Object,
			default: () => ({})
		}
	},
	computed: {
		...mapGetters(["user"])
	},
	data() {
		return {
			assetsBaseUrl: process.env.VUE_APP_PUBLIC_BASE_URL
		}
	}
}
</script>

<style lang="scss" scoped>
@import "../styles/_variables.scss";
.loading_ani,
.article {
	padding: 20px 20px 10px 20px;
    margin-bottom: 20px;
}
.article {
	margin-bottom: 10px;
	&:hover {
		.edit_article,
		.del_article {
			display: inline-block;
		}
	}
    .content {
		overflow: hidden;
    }
	.avatar {
		float: right;
		font-size: 0;
		img {
			width: 120px;
			height: 120px;
			border-radius: 4px;
			transition: all .6s ease;
			&:hover {
				transform: scale(1.11)
			}
		}
	}
    .abstract {
        height: 96px;
        line-height: 24px;
		margin-top: 5px;
		padding-right: 5px;
        overflow: hidden;
        color: $fontColor;
    }
    .title {
        position: relative;
        font-size: 18px;
        font-weight: 700;
        color: #555;
        transition: all ease-in-out .3s;
        &::before {
            position: absolute;
            content: '';
            width: 100%;
            height: 1px;
            bottom: 0;
            background-color: #337ab7;
            transform: scaleX(0);
            transform-origin: 50% 0;
            transition: transform .3s ease-in-out;
        }
        &:visited {
            color: $shallowColor;
        }
        &:hover {
            color: $highLightColor;
            &::before {
                transform: scale(1);
            }
        }
    }
    .base_info {
		display: flex;
		margin-top: 5px;
		&, a {
			color: $grayColor;
			&:not(.right) {
				font-size: 12px;
			}
		}
		.info_rt {
			flex: 0 0 50px;
		}
		.info_lf {
			flex: 1;
		}
	}
	.edit_article,
	.del_article {
		margin-left: 5px;
		font-size: 12px;
		display: none;
	}
    .heart_num,
    .comment_num {
        margin-left: 10px;
        font-size: 14px;
        cursor: pointer;
        &::before {
			margin-right: 2px;
			vertical-align: middle;
        }
	}
	.heart_num::before {
		font-size: 15px;
	}
    .heart_num,
    .comment_num,
    .alias_update {
        transition: color .3s ease-in-out;
        &:hover,
        a:hover {
            color: $fontColor;
		}
		&.heart_num:hover {
			color: $pinkColor;
		}
    }
}

@media screen and (max-width: $screenMinWidth) {
	.article .avatar img {
		margin-top: 30px;
		width: 90px;
		height: 90px;
	}
}
</style>
