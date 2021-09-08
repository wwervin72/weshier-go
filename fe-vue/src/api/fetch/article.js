import request from '../request'

export function ArticleFetch(params) {
	return request({
		baseURL: process.env.VUE_APP_BASE_API + '/article',
		...params
	})
}

/**
 * 创建文章
 * @param {*} data
 */
export function createArticle(data) {
	return ArticleFetch({
		url: '',
		method: 'post',
		data
	})
}

/**
 * 更新文章
 * @param {*} data
 */
export function updateArticle(data) {
	return ArticleFetch({
		url: '',
		method: 'put',
		data
	})
}

/**
 * 获取文章详情
 */
export function fetchArticleInfo(id) {
	return ArticleFetch({
		url: `/detail/${id}`,
	})
}

/**
 * 文章点赞
 */
export function HeartArticle(id) {
	return ArticleFetch({
		url: `/heart/${id}`,
	})
}

/**
 * 取消文章点赞
 */
export function CancelHeartArticle(id) {
	return ArticleFetch({
		url: `/heart/${id}`,
		method: "delete"
	})
}

/**
 * article 分页查询
 */
export function articlePagination(params) {
	return ArticleFetch({
		url: `/page`,
		params
	})
}
