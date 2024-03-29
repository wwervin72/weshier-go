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
 * 喜欢留言
 * @param {*} commentId
 */
export function HeartComment(commentId) {
	return CommentFetch({
		url: `/heart/${commentId}`
	})
}

/**
 * 取消喜欢留言
 * @param {*} commentId
 */
export function CancelHeartComment(commentId) {
	return CommentFetch({
		url: `/heart/${commentId}`,
		method: 'delete'
	})
}

/**
 * 删除留言
 * @param {*} data
 */
export function DeleteComment(data) {
	return CommentFetch({
		url: '',
		method: 'delete',
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
