import request from '../request'

export function CommentFetch(params) {
	return request({
		baseURL: process.env.VUE_APP_BASE_API + '/comment',
		...params
	})
}

/**
 * 留言
 * @param {*} data
 */
export function LeaveComment(data) {
	return CommentFetch({
		url: '',
		method: 'post',
		data
	})
}

/**
 * 评论文章
 * @param {*} data
 */
export function CommentArticle(data) {
	return CommentFetch({
		url: '/article',
		method: 'post',
		data
	})
}

/**
 * 获取留言分页
 * @param {*} data
 */
export function commentPagination(params) {
	return CommentFetch({
		url: '/page',
		params
	})
}

/**
 * 获取留言分页
 * @param {*} data
 */
export function articleCommentPagination(articleId, params) {
	return CommentFetch({
		url: `/article/${articleId}/page`,
		params
	})
}
