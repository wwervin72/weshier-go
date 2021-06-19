import request from '../request'

export function TagFetch(params) {
	return request({
		baseURL: process.env.VUE_APP_BASE_API + '/tag',
		...params
	})
}

/**
 * 创建 tag
 * @param {*} data
 */
export function createTag(data) {
	return TagFetch({
		url: '',
		method: 'post',
		data
	})
}

/**
 * 更新 tag
 * @param {*} data
 */
export function updateTag(data) {
	return TagFetch({
		url: '',
		method: 'put',
		data
	})
}

/**
 * 获取 tag 详情
 */
export function fetchTagInfo(id) {
	return TagFetch({
		url: `/detail/${id}`,
	})
}

/**
 * 获取 tag 列表
 */
export function fetchTagList() {
	return TagFetch({
		url: `/list`,
	})
}

/**
 * 获取用户 tag 列表
 */
export function fetchUserTagList() {
	return TagFetch({
		url: `/user/list`,
	})
}

/**
 * 获取 tag 分页数据
 */
export function fetchTagPagination(params) {
	return TagFetch({
		url: `/page`,
		params
	})
}
