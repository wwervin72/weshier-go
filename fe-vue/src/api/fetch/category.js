import request from '../request'

export function CategoryFetch(params) {
	return request({
		baseURL: process.env.VUE_APP_BASE_API + '/cate',
		...params
	})
}

/**
 * 创建分类
 * @param {*} data
 */
export function createCategory(data) {
	return CategoryFetch({
		url: '',
		method: 'post',
		data
	})
}

/**
 * 获取分类详情
 */
export function fetchCategoryInfo(id) {
	return CategoryFetch({
		url: `/detail/${id}`,
	})
}

/**
 * 更新分类
 * @param {*} data
 */
 export function updateCategory(data) {
	return CategoryFetch({
		url: '',
		method: 'put',
		data
	})
}

/**
 * 获取分类详情
 */
export function fetchCategoryList() {
	return CategoryFetch({
		url: `/list`,
	})
}

/**
 * 获取用户分类详情
 */
export function fetchUserCategoryList() {
	return CategoryFetch({
		url: `/user/list`,
	})
}

/**
 * 获取分类分页
 */
export function fetchCategoryPagination(params) {
	return CategoryFetch({
		url: `/page`,
		params
	})
}

