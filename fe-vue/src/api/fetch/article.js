import request from '../request'

export function ArticleFetch(params) {
	return request({
		baseURL: process.env.VUE_APP_BASE_API + '/article',
		...params
	})
}

/**
 * 创建 tag
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
 * 创建 tag
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
 * 获取 tag 详情
 */
export function fetchArticleInfo(id) {
	return ArticleFetch({
		url: `/detail/${id}`,
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
